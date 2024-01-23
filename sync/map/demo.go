package main

import (
	"fmt"
	"sync"
	"time"
)

var TestMap map[string]string

func init() {
	TestMap = make(map[string]string, 1)
}

var lock sync.Mutex

func main() {
	for i := 0; i < 1000; i++ {
		go Write("aaa")
		go Read("aaa")

		go Write("bbb")
		go Read("bbb")
	}
	time.Sleep(3 * time.Second)
}
func Read(key string) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Println(TestMap[key])
}

func Write(key string) {
	lock.Lock()
	TestMap[key] = key
	lock.Unlock()
}

//报错 fatal error: concurrent map writes
