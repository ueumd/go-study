package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

/**
Go 1.8 版本开始，net/http原生支持使用http.ShutDown来优雅的关停HTTP 服务。
这种方案同样要求用户创建自定义的 http.Server 对象，因为Shutdown方法无法通过其它途径调用。

我们来看下面的代码，这段代码通过结合捕捉系统信号（Signal）、goroutine 和管道（Channel）来实现服务器的优雅停止
*/

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	// 创建系统信号接收器
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")

	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
