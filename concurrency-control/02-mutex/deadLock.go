package main

import (
	"fmt"
	"sync"
	"time"
)

/**
死锁

双方都持有自己的资源，而去要求对方的资源。而且自己的资源自己持有，不可剥夺，必然产生死锁。
*/

func main() {
	// 派出所证明
	var psCertificate sync.Mutex

	// 物业证明
	var propertyCertificate sync.Mutex

	var wg sync.WaitGroup

	// 需要派出所和物业都处理
	wg.Add(2)

	// 派出所处理
	go func() {
		defer wg.Done()

		psCertificate.Lock()
		defer psCertificate.Unlock()

		//检查材料
		time.Sleep(5 * time.Second)

		// 请求物业证明
		propertyCertificate.Lock()
		propertyCertificate.Unlock()

	}()

	// 物业处理
	go func() {
		defer wg.Done()

		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		//检查材料
		time.Sleep(5 * time.Second)

		// 请求派出所证明
		psCertificate.Lock()
		psCertificate.Unlock()
	}()

	wg.Wait()
	fmt.Println("成功完成")
}
