package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum2    int = 0
	rwmutex sync.RWMutex
)

func main() {

	// 读多写少场景

	// 写操作
	for i := 0; i < 10; i++ {
		go write(10)
	}

	// 读操作
	for i := 0; i < 10; i++ {
		go readSum2()
		go readSum2()
		go readSum2()
	}

	time.Sleep(3 * time.Second)

}

func readSum2() {
	rwmutex.RLock()
	ret := sum2
	fmt.Printf("Read data: %v\n", ret)
	rwmutex.RUnlock()
}

func write(i int) {
	rwmutex.RLock()
	sum2 += i
	fmt.Printf("===================Write Data: %v %d \n", sum2, i)
	rwmutex.RUnlock()

}
