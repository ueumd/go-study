package main

import (
	"sync"
	"sync/atomic"
)

// 标记
var initialized uint32

var lock sync.Mutex

type singleton struct {
}

func (this *singleton) connection() {
	println("connection method")
}

var instance *singleton

func getInstance() *singleton {

	// 如果标记为被设置，直接返回，不加锁
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	/*
		代码虽然解决了线程安全，但是每次调用GetInstance()都要加锁会极大影响性能。
		所以接下来可以借助"sync/atomic"来进行内存的状态存留来做互斥
	*/

	//如果没有，则加锁申请
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)

		// 设置标记位
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

func main() {
	s := getInstance()
	s.connection()
}
