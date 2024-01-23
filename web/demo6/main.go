package main

import (
	"golang-study/web/demo5/ginga"
	"net/http"
)

func main() {
	server := ginga.New()

	// server.Use(ginga.Logger())

	server.GET("/", func(ctx *ginga.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index</h1>")
	})

	server.GET("/hello", func(ctx *ginga.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	server.GET("/panic", func(ctx *ginga.Context) {
		names := []string{"1", "2", "3"}
		ctx.String(http.StatusOK, names[100])
		// ctx.HTML(http.StatusOK, "<h1>panic</h1>")
	})

	server.Run(":5173")
}
