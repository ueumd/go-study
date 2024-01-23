package main

import (
	"fmt"
	"sync"
	"time"
)

/**
这段代码循环 1000 次，每次给 sum 加 10，理论上应该是 10000，但是执行结果为 8380、或者 9010 或者 9130 等。

原因是多个 go 语句并发地对 sum 进行加 10 操作，不能保证每次取的值就是上一次执行的结果。
*/

var sum int = 0
var mutex sync.Mutex

func main() {
	for i := 0; i < 1000; i++ {
		go add(10)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("sum is ", sum)

}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
