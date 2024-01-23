package main

import (
	"fmt"
	"net/http"
)

type TestRouter struct {
	content string
}

func (handler *TestRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, handler.content)
}

func main() {

	// 使用 http.HandleFunc 注册路由
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.PATH=%q\n", request.URL.Path)
	})

	http.HandleFunc("/hello", helloHandlerTest)

	http.Handle("/test", &TestRouter{"TestRouter"})

	// port
	http.ListenAndServe(":5173", nil)
}

func helloHandlerTest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PATH=%q\n", req.URL.Path)
}

// ----------------------------------------------------------
type myHandler interface {
	myServeHTTP(writer interface{}, request interface{})
}

// myHandle 参数 handler
func myHandle(pattern string, handler myHandler) {
	handler.myServeHTTP(pattern, handler)
}

type myTest struct {
	content string
}

func (self *myTest) myServeHTTP(writer interface{}, request interface{}) {
	fmt.Println(self.content)
}

func init() {
	myHandle("/", &myTest{"hello --------"})
}
