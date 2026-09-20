package gee

import "net/http"

type geeEngine struct {
	router      map[string]HandlerFunc
	middlewares []HandlerFunc
}

func New() *geeEngine {
	return &geeEngine{
		router:      make(map[string]HandlerFunc),
		middlewares: []HandlerFunc{},
	}
}

func (engine *geeEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v, o := engine.router[r.Method+"-"+r.URL.Path]
	if !o {
		http.NotFound(w, r)
		return
	}

	context := &Context{
		middlewares: engine.middlewares,
		index:       -1,
		W:           w,
		Req:         r,
		handler:     v,
	}

	context.Next()
}

func (engine *geeEngine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (gee *geeEngine) addRoute(method string, pattern string, handle HandlerFunc) {
	key := method + "-" + pattern
	gee.router[key] = handle
}

func (gee *geeEngine) Use(middlewares ...HandlerFunc) {
	gee.middlewares = append(gee.middlewares, middlewares...)
}

func (gee *geeEngine) GET(pattern string, handle HandlerFunc) {
	gee.addRoute("GET", pattern, handle)
}

func (gee *geeEngine) POST(pattern string, handle HandlerFunc) {
	gee.addRoute("POST", pattern, handle)
}
