package main

import (
	"fmt"
	"sync"
	"time"
)

var testMap map[string]string

/*
* 并发安全
Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。
不安全原因： 同一个变量在多个goroutine中访问需要保证其安全性。

原因：因为map变量为 指针类型变量，并发写时，多个协程同时操作一个内存，类似于多线程操作同一个资源会发生竞争关系，共享资源会遭到破坏
因此golang 出于安全的考虑，抛出致命错误：fatal error: concurrent map writes。

解决：在写操作的时候增加锁，删除时候除了加锁外，还需要增加断言避免出现错误
*/
var lock sync.RWMutex

func main() {
	// 开辟空间
	testMap = make(map[string]string)

	for i := 0; i < 1000; i++ {
		go write("aaa")
		go read("aaa")
		go write("bbb")
		go read("bbb")
	}

	time.Sleep(8 * time.Second)
}

func read(key string) {
	lock.Lock()
	fmt.Println(testMap[key])
	lock.Unlock()
}

func write(key string) {
	lock.Lock()
	testMap[key] = key
	lock.Unlock()
}
