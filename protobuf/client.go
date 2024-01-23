package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
	"ueumd/protobuf/proto"
)

func main()  {
	// 1. 连接
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())

	if err != nil{
		log.Printf("连接失败：[%v]", err)
	}

	// 2. 关闭连接
	defer conn.Close()

	// 3. 声明客户端
	client := proto.NewGreeterTestClient(conn)

	// 4. 发送数据
	rsp, _ := client.SayHello(context.Background(), &proto.HelloRequest{
		Name: "protobuf",
		Url: "github.com/ueumd/go-test/protobuf",
		Sex: proto.Gender_MALE,
		List: map[string]string{
			"name":"ueumd",
			"company": "XREDU",
		},
		AddTime: timestamppb.New(time.Now()),
	})

	fmt.Println("收到消息：", rsp.Message)
}