package main

import (
	"fmt"
	"sync"
)

/*
*
死锁 Copy已使用的Mutex
*/
type counter struct {
	sync.Mutex
	count uint64
}

func main() {
	var c counter

	c.Lock()
	defer c.Unlock()

	c.count++

	// 造成死锁
	// c.count++ 执行完后进行foo(c) 方法中执行
	// 在foo(c) 方法中又进行上锁
	foo(c)
}

func foo(c counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

/**
检测
go vet main.go

# command-line-arguments

.\main.go:24:6: call of foo copies lock value: command-line-arguments.counter
.\main.go:27:12: foo passes lock by value: command-line-arguments.counter

*/
