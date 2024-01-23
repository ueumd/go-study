package main

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

//type HandlerFunc interface {
//	ServeHTTP(http.ResponseWriter, *http.Request)
//}

type Engine struct {
	router map[string]HandlerFunc
}

func NewEngine() *Engine {
	return &Engine{
		// 开辟空间
		router: make(map[string]HandlerFunc),
	}
}

func (engine *Engine) addRoute(method string, patter string, handler HandlerFunc) {
	key := method + "-" + patter
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 实现ServeHTTP接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path

	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func main() {
	r := NewEngine()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.Run(":5173")
}
