package main

import (
	"fmt"
	"time"
)

func main() {

	chanStr := make(chan string)

	// 创建一个匿名协程
	go func() {
		// 发送数据：Hello Golang"
		chanStr <- "Hello Golang"

		time.Sleep(time.Millisecond)

		close(chanStr)

	}()

	for {
		// 接受数据:  Hello Golang
		msg, isOK := <-chanStr
		if isOK {
			// 打印
			fmt.Println("chan Msg =", msg) // chan Msg = Hello Golang
		} else {
			fmt.Println("通道已关闭")
			break
		}
	}

	/*
		chan Msg = Hello Golang
		通道已关闭
	*/

}
