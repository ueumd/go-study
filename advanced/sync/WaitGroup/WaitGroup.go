package main

import (
	"fmt"
	"sync"
	"time"
)

func p() {
	for {
		fmt.Println("I am p method")
	}
}

func main1() {
	//  子协程
	go p()

	// 主协程

	// 主死随从
	fmt.Println("I am main method")

	time.Sleep(time.Second * 2)
}

func print4(j int) {
	fmt.Println("I am p method", j)
}

func main2() {
	for i := 0; i < 10; i++ {
		go print4(i)
	}

	fmt.Println("I am main method")

	// 10 s
	time.Sleep(time.Second * 10)
	// 或者
	// 无限循环
	for {
	}
}

/**
解决主协程的goroutine 在子协程结束后自动结束

WaitGroup 3个函数
Add()
Done()

Wait()

Add 的数量 和 Done的数量必须相等
*/

func print(n int, wg *sync.WaitGroup) {
	defer wg.Done() // 减1
	fmt.Println("I am p method", n)
}

func main() {
	var wgg sync.WaitGroup

	wgg.Add(10) // 总共有多少个

	for i := 0; i < 10; i++ {
		go print(i, &wgg)
	}

	// 阻塞主协程
	wgg.Wait()

	fmt.Println("I am main method")
}
