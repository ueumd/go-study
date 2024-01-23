package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5173")
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")

		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		// 发送数据
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}

		// 接收服务端数据
		buf := [512]byte{}

		// 读取整个buf
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}

		// 从开始读取到n，不包括n
		fmt.Println(string(buf[:n]))
	}
}
