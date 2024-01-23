package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter 线程安全的计数器
type counter struct {
	mu    sync.Mutex
	count uint64
}

// Incr 对计数值加一
func (c *counter) incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Count 获取当前的计数值
func (c *counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// sleep 1秒，然后计数值加1
func worker(c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.incr()
}

func main() {
	var ct counter

	var wg sync.WaitGroup

	wg.Add(10) // WaitGroup的值设置为10

	for i := 0; i < 10; i++ { // 启动10个goroutine执行加1任务
		go worker(&ct, &wg)
	}

	// 检查点，等待goroutine都完成任务
	wg.Wait()
	// 输出当前计数器的值
	fmt.Println(ct.Count())
}
