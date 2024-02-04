package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(3)

	var channel <-chan interface{}
	var wg sync.WaitGroup
	const numGoroutines = 1000000 // 1M

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	blockFunc := func() {
		wg.Done()
		<-channel
	}

	wg.Add(numGoroutines)

	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go blockFunc()
	}

	wg.Wait()

	after := memConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1024)
}
