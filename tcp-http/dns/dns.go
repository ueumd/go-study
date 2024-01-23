package main

import (
	"fmt"
	"net"
)

func main() {
	// 域名改成自己要测试的
	dns := "chat8.hzqyi.cn"
	// 解析cname
	cname, _ := net.LookupCNAME(dns)
	fmt.Println("cname:", cname)

	// 解析ip地址
	ips, err := net.LookupHost(dns)
	if err != nil {
		fmt.Println("Err: ", err.Error())
		return
	}
	fmt.Println(ips)

	// 反向解析(主机必须得能解析到地址), IP地址改成你的
	dnsName, _ := net.LookupAddr("221.181.204.234")
	fmt.Println("Hostname:", dnsName)
}
