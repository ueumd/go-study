package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(greet)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func greet() {
	fmt.Println("hello world")
}
