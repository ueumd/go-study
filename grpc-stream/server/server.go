package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
	"ueumd/grpc-stream/proto"
)

const PORT =":8082"

type server struct {}

// 之前的写法 通过不了
//func (s *server)GetStream(ctx context.Context, req *proto.StreamReqData )(*proto.StreamResData, error)  {
//	return nil, nil
//}

/**
stream.pb.go

// GreeterServer is the server API for Greeter service.

type GreeterServer interface {
	GetStream(*StreamReqData, Greeter_GetStreamServer) error
	PutStream(Greeter_PutStreamServer) error
	AllStream(Greeter_AllStreamServer) error
}
 */

// 服务端模式
// 客户端发起一次请求，服务端返回一段连续的数据流
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer ) error  {
	i := 0
	for {
		i ++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})

		time.Sleep(time.Second)

		// 发送10次服务停止发送
		if i > 10 {
			break
		}
	}
	return nil
}

// 客户端流模式
// 与服务端数据流模式相反，客户端源源不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer ) error  {
	for {
		if res, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(res.Data)
		}
	}
	return nil
}

// 双向流模式
// 客户端和服务端都可以向对方发送数据流
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer ) error  {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			res, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + res.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{
				Data: "I am server",
			})
			time.Sleep(time.Second*2)
		}
	}()
	wg.Wait()
	return nil
}



func main()  {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	rpcServer := grpc.NewServer()

	proto.RegisterGreeterServer(rpcServer, &server{})

	err = rpcServer.Serve(lis)

	if err != nil {
		panic(err)
	}

}