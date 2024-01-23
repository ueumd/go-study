package main

import (
	"golang-study/web/demo2/ginga"
	"net/http"
)

func main() {
	r := ginga.New()

	r.GET("/", func(ctx *ginga.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(ctx *ginga.Context) {
		// hello?id=110
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("id"), ctx.Path)
	})

	r.POST("/login", func(ctx *ginga.Context) {
		ctx.JSON(http.StatusOK, ginga.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.Run(":5173")
}
