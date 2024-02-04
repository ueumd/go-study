package main

import (
	"fmt"
	"runtime"
	"sync"
)

// goroutine的竞争状态

var (
	counter int // 所有goroutine的共享变量
	wg      sync.WaitGroup
)

// 改变counter的函数
func add(id int) {
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()

	for i := 0; i < 2; i++ {
		// 捕获counter的值
		value := counter
		// 当前goroutine从线程中退出，重新放回到队列，切换其它线程
		runtime.Gosched()

		// 增加本地value的值
		value++
		// 将该值保存到counter
		counter = value
	}

}

func main() {
	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	go add(1)
	go add(2)

	// 等待goroutine执行结束
	wg.Wait()
	fmt.Println("Final counter:", counter)
}
