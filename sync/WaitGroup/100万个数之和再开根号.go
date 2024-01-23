package main

import (
	"fmt"
	"math"
	"sync"
)

// 计算1000个数的和
func compute(m *sync.Mutex, wg *sync.WaitGroup, s, e int, count *int) {
	sum := 0
	for i := s; i < e; i++ {
		sum += i
	}

	m.Lock()
	*count += sum
	m.Unlock()
	wg.Done()
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	var n int = 1000

	var count int
	wg.Add(n)

	for i := 0; i < n; i++ {
		go compute(&mutex, &wg, i*n+1, (i+1)*n+1, &count)
	}

	wg.Wait()
	fmt.Println(math.Sqrt(float64(count)))

	return
}
