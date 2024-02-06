package main

import "fmt"

// 接收通道的值
func recValue(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main2() {

	// 缓冲的通道必须有接收才能发送
	ch := make(chan int)

	// 如果没有goroutine 接收，这里会造成 deadlock
	go recValue(ch) // 启用goroutine从通道接收值

	ch <- 10
	fmt.Println("发送成功")
}

func main33() {
	c := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i = i + 1 {
	//		c <- i
	//	}
	//	close(c)
	//}()

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("Finished")
}

// 死锁4：读取空channel 死锁
func main() {
	// 死锁1
	ch := make(chan int)
	//close(ch) 向关闭的channel中读取数据 是该数据的类型的零值
	num := <-ch
	fmt.Println("num=", num)
}
