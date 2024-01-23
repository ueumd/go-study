package main

import (
	"google.golang.org/grpc"
	"net"
	"ueumd/grpc/server/service"
)

func main()  {
	rpcServer := grpc.NewServer()
	service.RegisterProdServiceServer(rpcServer, &service.ProdService{})
	// service.RegisterProdServiceServer(rpcServer, new(service.ProdService))


	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = rpcServer.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}