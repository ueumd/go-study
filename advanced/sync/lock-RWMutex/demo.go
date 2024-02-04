package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   int = 0
	mutex sync.Mutex
)

func main() {

	// 写操作
	for i := 0; i < 10; i++ {
		go add(10)
	}

	// 读操作
	for i := 0; i < 10; i++ {
		go readSum()
		go readSum()
		go readSum()
	}

	time.Sleep(2 * time.Second)

}

func readSum() {
	// 加锁保证读取到的是最终结果
	mutex.Lock()
	ret := sum
	fmt.Printf("Read data: %v\n", ret)
	mutex.Unlock()
}

func add(i int) {
	mutex.Lock()
	sum += i
	fmt.Printf("===================Write Data: %v %d \n", sum, i)
	mutex.Unlock()
}
