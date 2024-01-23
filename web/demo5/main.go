package main

import (
	"golang-study/web/demo5/ginga"
	"net/http"
)

func main() {
	server := ginga.New()

	server.Use(ginga.Logger())

	server.GET("/", func(ctx *ginga.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	v1 := server.Group("/v1")
	v1.GET("/", func(ctx *ginga.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	server.Run(":5173")
}
