package main

import (
	"sync"
)

/*
借助Once来实现单例模式的实现
*/

// 标记
var once sync.Once

type singleton2 struct {
}

func (this *singleton2) connection() {
	println("connection method")
}

var instance2 *singleton2

func getInstance2() *singleton {
	once.Do(func() {
		instance2 = new(singleton2)
	})

	return instance
}

func main() {
	s := getInstance2()
	s.connection()
}
