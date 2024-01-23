package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr()) // 192.168.1.2:54564

	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0]) // 192.168.1.2

}
