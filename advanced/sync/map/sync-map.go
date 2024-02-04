package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}

	m.Store("a", 11)
	m.Store("b", 22)
	m.Store("c", 33)
	m.Store("d", 44)

	//读取数据
	fmt.Println(m.Load("a")) // 11 true

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	/**
	a 11
	b 22
	c 33
	d 44
	*/
}
