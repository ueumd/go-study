package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ready = 0
var cond = sync.NewCond(&sync.Mutex{})
var count = 0

func main() {

	for i := 0; i < 10; i++ {
		go readyGo(i)
	}

	wake()

	fmt.Println("所有运动员都准备就绪，比赛开始。。。")
}

func readyGo(i int) {
	time.Sleep(time.Second * time.Duration(rand.Int63n(10)))
	// 加锁更改等待条件
	cond.L.Lock()
	ready++
	cond.L.Unlock()

	fmt.Printf("======运动员%d已准备就绪 \n", i)
	// 广播唤醒等待者，这里可以使用Broadcast和Signal
	cond.Signal()
}

func wake() {
	cond.L.Lock()

	for ready != 10 {
		cond.Wait()
		count++
		fmt.Printf("裁判员被唤醒%d次 \n", count)
	}

	cond.L.Unlock()
}
