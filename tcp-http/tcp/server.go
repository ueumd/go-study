package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn) {

	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte

		n, err := reader.Read(buf[:])

		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到 client端发来的数据：", recvStr)

		// 服务端发送数据
		conn.Write([]byte("server: " + recvStr))
	}
}

func main() {

	// 监听某个端口的tcp网络
	listen, err := net.Listen("tcp", "127.0.0.1:5173")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		// 等待下次请求过来并建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		// 在这个连接上做一些事情
		go process(conn)
	}
}
