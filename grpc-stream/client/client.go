package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
	"ueumd/grpc-stream/proto"
)

func main()  {

	// 添加grpc.WithInsecure()，不然没有证书会报错
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()


	client := proto.NewGreeterClient(conn)

	//服务端流模式
	res, _ := client.GetStream(context.Background(), &proto.StreamReqData{Data:"Golang Server GetStream"})
	for {
		// socket send recv
		result, err := res.Recv()

		// 服务端发送10次后停止发送，客户端会收到EOF
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("Data:", result.Data)
	}

	// 客户端模式
	putS, _ := client.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("Golang Server PutStream %d", i),
		})

		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向
	allStr, _ := client.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			res, _ := allStr.Recv()
			fmt.Println("收到服户端消息：" + res.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{
				Data: "Golang Server AllStream ",
			})
			time.Sleep(time.Second*2)
		}
	}()
	wg.Wait()

}
