package gee

import "net/http"

type GeeEngine struct {
	router map[string]HandlerFunc

	middlewares []HandlerFunc
}

func New() *GeeEngine {
	return &GeeEngine{
		router:      make(map[string]HandlerFunc),
		middlewares: []HandlerFunc{},
	}
}

func (engine *GeeEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v, o := engine.router[r.Method+"-"+r.URL.Path]
	if !o {
		http.NotFound(w, r)
		return
	}

	context := NewContext(w, r, v)
	v(context)
}

func (engine *GeeEngine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (gee *GeeEngine) addRoute(method string, pattern string, handle HandlerFunc) {
	key := method + "-" + pattern
	gee.router[key] = handle
}

func (gee *GeeEngine) Use(middlewares ...HandlerFunc) {
	gee.middlewares = append(gee.middlewares, middlewares...)
}

func (gee *GeeEngine) GET(pattern string, handle HandlerFunc) {
	gee.addRoute("GET", pattern, handle)
}

func (gee *GeeEngine) POST(pattern string, handle HandlerFunc) {
	gee.addRoute("POST", pattern, handle)
}
