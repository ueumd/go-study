package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func welcome(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("welcome"))
}

type HelloHandle struct {
	content string
}

type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*
*
第一种方式 HandleFunc
*/
func serveHttp1() {
	http.HandleFunc("/", welcome)

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {

		writer.WriteHeader(401)

		res := HttpResponse{Code: 401, Msg: "SUCCESS", Data: request.RemoteAddr}

		jsonRes, err := json.Marshal(res)
		if err != nil {
			fmt.Fprintln(writer, err.Error())
			return
		}

		writer.Write(jsonRes)

	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome teachers!")
	})

	http.ListenAndServe(":8080", nil)
}

/*
*
自定义的 http.Handler 接口的实例
*/
func (handler *HelloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.content)
}

/*
*
第二种方式 Handle
*/
func serveHttp2() {
	http.Handle("/hello", &HelloHandle{content: "Hello World"})

	http.ListenAndServe(":8080", nil)
}

/**

http.HandleFunc 和 http.Handle都是用于注册路由，
前者是一个具有func(w http.ResponseWriter, r *http.Requests)签名的函数
而后者是一个结构体，该结构体实现了func(w http.ResponseWriter, r *http.Requests)签名的方法。
*/

/*
*
第三种方式
ServeMux(服务复用器)
*/
func serveHttp3() {
	//声明多路复用mux对象
	router := http.NewServeMux()
	router.HandleFunc("/three", func(writer http.ResponseWriter, request *http.Request) {
		write, err := writer.Write([]byte("ServeMux(服务复用器)"))
		if err != nil {
			return
		}
		fmt.Println(write)
	})
	//通过实现mux的ServeHTTP方法可实现路由功能
	http.ListenAndServe("127.0.0.1:8080", router) //路由注册
}

/*
*
第四种方式
server
*/
func serveHttp4() {

	mux := http.NewServeMux()

	mux.HandleFunc("/four", func(writer http.ResponseWriter, request *http.Request) {
		write, err := writer.Write([]byte("four"))
		if err != nil {
			return
		}
		fmt.Println(write)
	})

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Handler:      mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func main() {
	serveHttp1()
	// serveHttp2()
	// serveHttp3()
	// serveHttp4()
}
