package main

import "fmt"

// fatal error: all goroutines are asleep - deadlock!
func maindeadlock() {
	var done = make(chan bool, 1)

	// 造成死锁
	done <- true
	fmt.Println(<-done)
}

func main1() {
	var done = make(chan bool, 1)

	// 造成死锁
	done <- true
	// fmt.Println(<-done)

	// 解决方式1
	println(<-done)
}

// 解决方式2

// 方式1
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10

	fmt.Println("发送成功")
	for {
	}
}
