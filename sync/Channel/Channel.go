package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func consumer(queue chan int)  {
	defer wg.Done()
	data := <-queue
	fmt.Println(data)
}

func consumer2(queue chan int)  {
	defer wg.Done()
	// 这样取没法判断 channel是否关闭
	for data := range queue {
		fmt.Println(data)
	}
}

func consumer3(queue chan int)  {
	defer wg.Done()
	for {
		// 通过ok 来确定 channel是否关闭
		data, ok := <- queue
		if !ok {
			// 不终止会一直打印0
			break
		}
		fmt.Println(data)
	}
}

/**
channel 要配合 goroutine 有写有读 才不会造成死锁
 */
func main()  {

	var msg  chan int

	// go中使用make初始化的有三种： 1. slice 2.map 3.channel

	// 无缓冲方式 初始化 要注意死锁问题  deadlock
	// fatal error: all goroutines are asleep - deadlock!

	// msg = make(chan int)

	// 有缓冲方式
	msg = make(chan int, 1)

	// 将1放入到msg channel中
	msg <- 100

	wg.Add(1)
	// go consumer(msg)
	// go consumer2(msg)

	go consumer3(msg)

	msg <- 200
	msg <-300

	// 关闭channel
	// 1. 已经关闭的channel 不能 再送发数据
	// 2. 已经关闭的channel 可以 读取数据，直到数据取完为止
	close(msg)

	// msg <- 300 // panic
	wg.Wait()

}