package ginga

import (
	"log"
	"net/http"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handlers[key]; ok {
		// handler(ctx.Writer, ctx.Req)
		handler(ctx)
	} else {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
