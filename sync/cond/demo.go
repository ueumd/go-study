package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁需要保护的条件变量
var done = false

// 调用 Wait() 等待通知，直到 done 为 true。
func read(name string, c *sync.Cond) {
	c.L.Lock()

	for !done {
		// 调用wait
		c.Wait()
	}

	fmt.Println(name, "--------------- starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts writing")

	// 暂停1s
	time.Sleep(time.Second)

	c.L.Lock()
	// 将 done 置为 true
	done = true
	c.L.Unlock()

	fmt.Println(name, "wakes all")
	// Broadcast
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)

	write("writer", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}
