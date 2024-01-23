package main

import (
	"fmt"
	"golang-study/web/demo3/ginga"
	"net/http"
)

func main() {
	server := ginga.New()

	//r.GET("/", func(ctx *ginga.Context) {
	//	ctx.HTML(http.StatusOK, "<h1>Hello</h1>")
	//})

	server.GET("/hello/:id", func(ctx *ginga.Context) {
		// hello?id=110
		fmt.Println("xxx")
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("id"), ctx.Path)
	})

	server.POST("/login", nil)

	server.GET("/req/:id/:name", func(ctx *ginga.Context) {
		// hello?id=110
		ctx.String(http.StatusOK, "req %s, you're at %s\n", ctx.Query("id"), ctx.Path)
	})

	//server.POST("/login", func(ctx *ginga.Context) {
	//	ctx.JSON(http.StatusOK, ginga.H{
	//		"username": ctx.PostForm("username"),
	//		"password": ctx.PostForm("password"),
	//	})
	//})

	server.GET("/assets/*filepath", func(ctx *ginga.Context) {
		ctx.JSON(http.StatusOK, ginga.H{"filepath": ctx.Param("filepath")})
	})

	server.Run(":5173")
}
