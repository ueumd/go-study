package main

import (
	"fmt"
	"sync"
	"time"
)

var wgp2 sync.WaitGroup

// 2. 利用 chan
var stop2 chan bool = make(chan bool)

func cpuInfo2()  {
	defer wgp2.Done()
	for {
		select {
		case <- stop2:
			fmt.Println("退出cpu监控")
			return
		default:
			// 上面没有就执行下面
			time.Sleep(time.Second*2)
			fmt.Println("cpu信息读取完成")
		}

	}
}

func main()  {
	wgp2.Add(1)
	go cpuInfo2()

	time.Sleep(time.Second*6)
	// 5秒后不再监控
	stop2 <- true

	wgp2.Wait()
	fmt.Println("信息监控完成")
}