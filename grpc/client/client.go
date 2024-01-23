package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"ueumd/grpc/server/service"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8082端口
	// 并且添加grpc.WithInsecure()，不然没有证书会报错
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	client := service.NewProdServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	resp, err := client.GetProductStock(context.Background(), &service.ProductRequest{ProdId: 777})

	//ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	//resp, er := client.GetProductStock(ctx, &service.ProductRequest{ProdId: 777})
	//
	//if er != nil {
	//	st, ok := status.FromError(er)
	//
	//	if !ok{
	//		panic("解析error失败")
	//	}
	//	fmt.Println(st.Message())
	//	fmt.Println(st.Code())
	//}

	if err != nil {
		panic("调用gRPC方法错误: "  + err.Error())
	}

	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock) // 调用gRPC方法成功，ProdStock =  777
}
