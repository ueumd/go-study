package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 与本地的NC TCP服务器建立连接
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// 设置一个定时函数，三秒后关闭连接
	timer := time.AfterFunc(3*time.Second, func() {
		conn.Close()
	})

	// 读网络数据，三秒后Read还未成功，连接就会被关闭
	// 如果数据读取超时，Read就会返回use of closed network connection的错误
	// 如果数据读取成功，我们就关闭定时器，不关闭这个连接
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	timer.Stop()
	fmt.Printf("读取到: %q\n", string(buf[:n]))
}
