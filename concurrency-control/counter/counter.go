package main

import (
	"fmt"
	"sync"
)

func main2() {
	var count = 0

	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				// 每次打印结果都是不一样的
				// count不是一个原子操作
				// count++ 是一个临界区
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println("count:", count) // 期望值是 1000000，但是每次都不是，且结果不一样
}

func main3() {
	var count = 0

	var wg sync.WaitGroup

	var mu sync.Mutex

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				// 每次打印结果都是不一样的
				// count不是一个原子操作

				// 进入临界区锁上
				mu.Lock()
				count++
				// 执行完后解锁
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println("count:", count) // 期望值是 1000000
}

/**
上面的方法可以有效解决了问题，但是业务逻辑 锁的逻辑都暴露出去了
*/

// 线程安全的计数器类型
type counter struct {
	mu    sync.RWMutex
	count uint64
}

// +1 内部使用互斥锁保护
func (self *counter) incr() {
	self.mu.Lock()
	self.count++
	self.mu.Unlock()
}

// 得到计数器的值， 需要保护
func (self *counter) getCount() uint64 {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.count
}

func main() {
	var counter counter
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				// 直接调用封装好的方法
				counter.incr()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println("count:", counter.getCount()) // 期望值是 1000000
}
