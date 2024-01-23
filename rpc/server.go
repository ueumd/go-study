package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {}

// 注意方法名大写
func (s *HelloService) Hello(request string, reply *string) error  {
	*reply = "hello, " + request
	return nil
}

func main()  {
	// 1.实例一个server
	listener, _ := net.Listen("tcp", ":8081")


	// 2.注册逻辑
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 3. 启动服务
	for {
		conn, _:= listener.Accept() // 一个新连接进来时
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

