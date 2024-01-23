package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 把一个 IPv4 或者 IPv6 的地址转化成 IP 类型
	name := "192.168.10.105"
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}

	os.Exit(0)
}
