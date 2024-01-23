package main

import (
	"fmt"
	"time"
)

/**
关闭 一个通道意味着不能再向这个通道发送值了。这个特性可以 用来给这个通道的接收方传达工作已经完成的信息。
*/

func closeChan() {
	jobs := make(chan int, 5) //申请通道
	done := make(chan bool)

	go func() {
		for {
			//job 要接受的数据存放的变量
			//ok 表示管道是否关闭，如果为 false，则表明管道已经关闭
			job, ok := <-jobs
			if ok {
				fmt.Println("received job", job)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job: ", j)
	}

	time.Sleep(time.Second * 2)

	close(jobs)

	fmt.Println("sent all jobs")

	<-done //阻塞
}

func rangChan() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "tow"

	close(queue)
	fmt.Println(queue) //0xc04208a000
	for v := range queue {
		fmt.Println(v)
	}
}

func main() {
	closeChan()
	// rangChan()
}
