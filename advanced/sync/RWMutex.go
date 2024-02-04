package main

import (
	"fmt"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

var rwLock sync.RWMutex // 读写锁

/**
多个协程之间 读与读不会产生影响
 */
func read()  {
	defer wg3.Done()

	rwLock.RLock() // 读锁
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取成功")
	rwLock.RUnlock()
}

func write()  {
	defer wg3.Done()

	rwLock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second*10)
	fmt.Println("修改成功")
	rwLock.Unlock()
}

func main()  {

	wg3.Add(5)


	for i:=0; i< 5; i++ {
		go read()
	}

	for i:=0; i< 1; i++ {
		go write()
	}

	wg3.Wait()

}
