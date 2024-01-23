package main

import (
	"fmt"
	"sync"
)

/**
重入锁
Mutex 不是重入锁

以下列子 报错

实现 可重锁  goroutine方案 和 token方案
*/

func foo2(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock() // 第一次 加锁
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock() // 第二次 加锁
	fmt.Println("in bar")
	l.Unlock()
}

func main() {
	mu := &sync.Mutex{}

	foo2(mu)
}
