package main

import (
	"fmt"
	"sync"
)

var total int
var wgp sync.WaitGroup
var lock sync.Mutex  // 互斥锁

/**
互斥锁 读写锁 同步数据 能不用就不用 性能问题 耗时时间， 具体问题，具体分析
绝大多数系统 都是读多写少
 */

func add()  {
	defer wgp.Done()
	for i:=0; i< 100000; i++ {
		lock.Lock()
		total += 1
		lock.Unlock()
	}
}

func sub()  {
	defer wgp.Done()
	for i:=0; i< 100000; i++ {
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}


func main()  {
	wgp.Add(2)

	/**
	并发执行时 不能保证 total += 1  total -= 1 一个一个的执行 得加锁
	*/
	go add()
	go sub()

	wgp.Wait()

	fmt.Println(total)  // 0 用上锁
}