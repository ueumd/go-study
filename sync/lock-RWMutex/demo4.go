package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

var Password = secret{password: "myPassword"}

// 通过rwmutex写
func rwChange(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("Change with rwmutex lock")
	time.Sleep(3 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

// 通过rwmutex读
func rwMutexShow(c *secret) string {
	c.RWM.RLock()
	fmt.Println("show with rwmutex", time.Now().Second())
	time.Sleep(1 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

// 通过mutex读，和rwMutexShow的唯一区别在于锁的方式不同
func mutexShow(c *secret) string {
	c.M.Lock()
	fmt.Println("show with mutex:", time.Now().Second())
	time.Sleep(1 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func main() {
	// 定义一个稍后用于覆盖(重写)的函数
	var show = func(c *secret) string { return "" }

	// 通过变量赋值的方式，选择并重写showFunc函数
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!", time.Now().Second())
		show = rwMutexShow
	} else {
		fmt.Println("Using sync.Mutex!", time.Now().Second())
		show = mutexShow
	}

	var wg sync.WaitGroup

	// 激活5个goroutine，每个goroutine都查看
	// 根据选择的函数不同，showFunc()加锁的方式不同
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", show(&Password), time.Now().Second())
		}()
	}

	// 激活一个申请写锁的goroutine
	go func() {
		wg.Add(1)
		defer wg.Done()
		rwChange(&Password, "123456")
	}()
	// 阻塞，直到所有wg.Done
	wg.Wait()
}
