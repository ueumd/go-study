package ginga

import (
	"fmt"
	"net/http"
	"strings"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    map[string]*node{},
		handlers: make(map[string]HandlerFunc),
	}
}

//func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
//	log.Printf("Route %4s - %s", method, pattern)
//	key := method + "-" + pattern
//	r.handlers[key] = handler
//}

//func (r *router) handle(ctx *Context) {
//	key := ctx.Method + "-" + ctx.Path
//	if handler, ok := r.handlers[key]; ok {
//		handler(ctx)
//	} else {
//		ctx.Writer.WriteHeader(http.StatusNotFound)
//		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
//	}
//}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
	fmt.Println(r.roots)
}

/*
*
注册的路由回调函数执行
*/
func (r *router) handle(ctx *Context) {
	n, params := r.getRoute(ctx.Method, ctx.Path)

	if n != nil {
		ctx.Params = params
		key := ctx.Method + "-" + n.pattern

		// 方法体执行
		// server.GET("/hello/:id", func(ctx *ginga.Context)
		r.handlers[key](ctx)
	} else {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

/*
	解析路由

/hello/:id
/req/:id/:name

parts [hello :id]
parts [req :id :name]
parts [assets *filepath]
*/
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	fmt.Println("parts", parts)
	return parts
}
