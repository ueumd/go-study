package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"ueumd/protobuf/proto"
)

const PORT =":8082"

type server struct {

}
func (ps *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: fmt.Sprintf("Name: %s, Url: %s, List: %s", request.Name, request.Url, request.List),
	}, nil
}


func main()  {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterGreeterTestServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)

	if err != nil {
		panic(err)
	}
}