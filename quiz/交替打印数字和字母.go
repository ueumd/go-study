package main

import (
	"fmt"
	"strings"
	"sync"
)

/**
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728


1. 两个 goroutine
2. 两个 channel 传递消息
3. 一个WaitGroup Add Wait Down
*/

func main() {
	number, letter := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Printf("%d%d", i, i+1)
				i += 2
				letter <- true
				// break
			}
		}
	}()

	// 加1
	wg.Add(1)

	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					// 减1
					// 读取完毕
					wait.Done()
					return
				}

				fmt.Print(str[i : i+2])

				i += 2

				// 打印数字
				number <- true

			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}
