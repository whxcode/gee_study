package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
	groups   map[string]*RouterGroup
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
		groups:   make(map[string]*RouterGroup),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0, len(vs))

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)

			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc, group *RouterGroup) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern

	_, ok := r.roots[method]

	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
	r.groups[key] = group
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	parmas := map[string]string{}
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)

		for index, part := range parts {
			if part[0] == ':' {
				parmas[part[1:]] = searchParts[index]
			}

			if part[0] == '*' && len(part) > 1 {
				parmas[part[1:]] = strings.Join(searchParts[index:], "/")

				break
			}
		}

		return n, parmas
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	n, parmas := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Parmas = parmas
		key := c.Method + "-" + n.pattern

		r.handlers[key](c)
		group := r.groups[key]

		if group != nil {
			middlewares := group.middlewares[:]

			for p := group.parent; p != nil; {
				middlewares = append(p.middlewares, middlewares...)
				p = p.parent
			}

			middlewares = append(middlewares, r.handlers[key])
			c.middlewares = middlewares
			c.Next()
		}

	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
