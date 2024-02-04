package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wgp4 sync.WaitGroup

/**
为了解决 Context1 和 Context2中的问题 go 1.7 提供Context.WithCancel
使用上变得更加优雅
*/

func cpuInfo4(ctx context.Context) {
	defer wgp4.Done()

	// go memoryInfo3(ctx)

	// ctx2, _ := context.WithCancel(context.Background()) // 一个新的ctx

	ctx2, _ := context.WithCancel(ctx) // 父子关系 父退出，子也会退出 链式取消 web开发中常用
	// 深层嵌套
	go memoryInfo4(ctx2)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			// 上面没有就执行下面
			time.Sleep(time.Second * 4)
			fmt.Println("cpu信息读取完成")
		}

	}
}

func memoryInfo4(ctx context.Context) {
	defer wgp4.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出内存监控")
			return
		default:
			// 上面没有就执行下面
			time.Sleep(time.Second * 2)
			fmt.Println("内存信息读取完成")
		}

	}
}

func main() {
	wgp4.Add(2)

	// 3s后超时后自动退出
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	go cpuInfo4(ctx)

	time.Sleep(time.Second * 1)

	// 6秒后不再监控
	// cancel() 超时后会自动取消

	wgp4.Wait()
	fmt.Println("信息监控完成")
}
