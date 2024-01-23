package main

import (
	"golang-study/web/demo4/ginga"
	"net/http"
)

func main() {
	server := ginga.New()

	server.GET("/", func(ctx *ginga.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	server.GET("/index", func(ctx *ginga.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := server.Group("/v1")
	v1.GET("/", func(ctx *ginga.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	v1.GET("/hello/:id", func(ctx *ginga.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("id"), ctx.Path)
	})

	v2 := server.Group("/v2")

	v2.GET("/login", nil)

	v2.GET("/req/:id/:name", func(ctx *ginga.Context) {
		// hello?id=110
		ctx.String(http.StatusOK, "req %s, you're at %s\n", ctx.Query("id"), ctx.Path)
	})

	server.Run(":5173")
}
