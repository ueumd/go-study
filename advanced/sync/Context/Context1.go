package main

import (
	"fmt"
	"sync"
	"time"
)

var wgp sync.WaitGroup

// 1. 利用全变量完成
var stop bool

func cpuInfo()  {
	defer wgp.Done()
	for {
		if stop {
			fmt.Println("退出cpu监控")
			break
		}
		time.Sleep(time.Second*2)
		fmt.Println("cpu信息读取完成")
	}
}

func main()  {
	wgp.Add(1)
	go cpuInfo()

	time.Sleep(time.Second*5)
	// 5秒后不再监控
	stop = true

	wgp.Wait()
	fmt.Println("信息监控完成")
}