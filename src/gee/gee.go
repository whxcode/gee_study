package gee

import "net/http"

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}

	engine.RouterGroup = &RouterGroup{
		prefix:      "",
		middlewares: []HandlerFunc{},
		parent:      nil,
		engine:      engine,
	}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix:      group.prefix + prefix,
		parent:      group,
		engine:      group.engine,
		middlewares: []HandlerFunc{},
	}

	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	engine.router.handle(context)
}

func (group *RouterGroup) addRoute(method string, pattern string, handle HandlerFunc) {
	_pattern := group.prefix + pattern
	group.engine.router.addRoute(method, _pattern, handle, group)
}

func (group *RouterGroup) GET(pattern string, handle HandlerFunc) {
	group.addRoute("GET", pattern, handle)
}

func (group *RouterGroup) POST(pattern string, handle HandlerFunc) {
	group.addRoute("POST", pattern, handle)
}

// ---
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
