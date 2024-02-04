package main

import (
	"fmt"
	"time"
)

func init() {
	/**
	 用 make 初始化通道
	chanName = make(chan chanType, bufferSize) 第二个参数是可选的，用于指定通道最多可以缓存多少个元素，默认值是 0，无缓冲通道

	court := make(chan int)
	*/

	//没有缓冲
	messages := make(chan string)
	//go func() { messages <- "ping" }()
	go func() {
		messages <- "ping" //匿名往通道 messages里发送一个消息ping
		time.Sleep(time.Millisecond)
		close(messages)
	}()

	for {
		msg, ok := <-messages //接收messages里的消息
		if !ok {
			fmt.Println("通道已关闭", ok)
			break
		}
		fmt.Println(msg, ok)
	}

}

/*
*

默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan） 通道准备好接收时，才允许进行发送（chan <-）。
可缓存通道 允许在没有对应接收方的情况下，缓存限定数量的值。
*/
func buff() {
	messages := make(chan string, 2) // make 了一个通道，最多允许缓存 2 个值。

	//因为这个通道是有缓冲区的，即使没有一个对应的并发接收 方，我们仍然可以发送这些值
	messages <- "buffered"
	messages <- "messages"

	//messages <- "messages111" // error
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/**
通道同步
使用通道来同步 Go 协程间的执行状态。
这里是一个 使用阻塞的接受方式来等待一个 Go 协程的运行结束。
*/
//通道作为函数的参数
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 3)
	fmt.Println("done")
	done <- true //向通道里发送消息
}
func testWorker() {
	done := make(chan bool, 1)
	go worker(done)

	/**
	如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker 还没开始运行时就结束了。
	*/
	<-done //程序将在接收到通道中 worker 发出的通知前一直阻塞
}

/**
通道方向
当使用通道作为函数的参数时，
你可以指定这个通道是不是 只用来发送或者接收值。
这个特性提升了程序的类型安全性
*/

// 一个只允许发送数据的通道
func ping(pings chan<- string, msg string) {
	// 写入数据
	pings <- msg
}

// pong 函数允许通道（pings）来接收数据，另一通道 （pongs）来发送数据。
func pong(pings <-chan string, pongs chan<- string) {
	// 读取pings数据
	msg := <-pings

	// 数据写入pongs
	pongs <- msg
}

func pingPong() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message") //向通道pings发送passed message消息
	pong(pings, pongs)            //接收passed message 又发送到通道pongs中

	// 再从pongs中读取数据
	fmt.Println("msg:", <-pongs) //打印接收消息
}

func main() {
	pingPong()
	for {
	}
	// testWorker()
}
