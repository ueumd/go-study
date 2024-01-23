package main

import (
	"fmt"
	"sync"
)

func main11() {
	var msg chan int

	// 无缓冲
	msg = make(chan int)

	msg <- 1 // 这里会阻塞，后面的代码也不会执行

	fmt.Println("deadlock")

	// 运行后直接报 fatal error: all goroutines are asleep - deadlock!
}

/**
为什么会报 deadlock？
1. 	msg <- 1 当向无缓冲channel msg写入数据时，会阻塞，阻塞之前会获取一把锁，锁的释话要等到channel 读取数据。如果有缓冲，数据先放入缓冲中，不会出现死锁
2. 使用 goroutine 来读取数据
3. 要有关闭锁
*/

func consumer33(msg chan int) {
	defer wg4.Done()
	fmt.Println(<-msg)
}

var wg4 sync.WaitGroup

func main() {
	var msg chan int
	msg = make(chan int)

	wg4.Add(1)
	go consumer33(msg)

	msg <- 1 // 这里会阻塞，所以 go consumer3(msg) 得放在上面

	fmt.Println("deadlock")
	wg4.Wait()

	// 运行后直接报 fatal error: all goroutines are asleep - deadlock!

}
