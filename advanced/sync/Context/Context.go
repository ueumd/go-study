package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wgp3 sync.WaitGroup

/*
*
为了解决 Context1 和 Context2中的问题 go 1.7 提供Context.WithCancel
使用上变得更加优雅
*/
func cpuInfo3(ctx context.Context) {
	defer wgp3.Done()

	// go memoryInfo3(ctx)

	// ctx2, _ := context.WithCancel(context.Background()) // 一个新的ctx

	ctx2, _ := context.WithCancel(ctx) // 父子关系 父退出，子也会退出 链式取消 web开发中常用
	// 深层嵌套
	go memoryInfo3(ctx2)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			// 上面没有就执行下面
			time.Sleep(time.Second * 2)
			fmt.Println("cpu信息读取完成")
		}
	}
}

func memoryInfo3(ctx context.Context) {
	defer wgp3.Done()
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
	wgp3.Add(2)

	// Background	返回一个空的 context，常作为根 context

	// WithCancel	基于父 context，生成一个可以取消的 context
	// 返回一个空的ctx
	ctx, cancel := context.WithCancel(context.Background())

	go cpuInfo3(ctx)
	// go memoryInfo3(ctx)

	time.Sleep(time.Second * 6)
	// 6秒后不再监控
	cancel()

	wgp3.Wait()
	fmt.Println("信息监控完成")
}
