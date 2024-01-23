package ginga

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

// 路由分组
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

// 分组
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// 定义Use函数，将中间件应用到某个 Group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

//func (engine *Engine) addRoute2(method string, pattern string, handler HandlerFunc) {
//	engine.router.addRoute(method, pattern, handler)
//}
//
//func (engine *Engine) GET(pattern string, handler HandlerFunc) {
//	engine.addRoute("GET", pattern, handler)
//}
//
//func (engine *Engine) POST(pattern string, handler HandlerFunc) {
//	engine.addRoute("POST", pattern, handler)
//}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

/*
*
实现ServeHTTP接口
*/
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var middlewares []HandlerFunc

	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	ctx := newContext(w, req)
	ctx.handlers = middlewares

	// 路由handle执行
	engine.router.handle(ctx)
}
