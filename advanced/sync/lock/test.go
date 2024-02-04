package main

import (
	"fmt"
	"sync"
	"time"
)

const MAXNUM = 1000 //map的大小
const LOCKNUM = 1e5 //加锁次数

var lock sync.Mutex        //互斥锁
var rwlock sync.RWMutex    //读写锁
var lock_map map[int]int   //互斥锁map
var rwlock_map map[int]int //读写锁map

func main() {
	var lock_w sync.WaitGroup
	var rwlock_w sync.WaitGroup
	lock_w.Add(LOCKNUM)
	rwlock_w.Add(LOCKNUM)
	lock_ch := make(chan int, 1000) // 缓存影响小，因为存入马上便从chan中取出来
	rwlock_ch := make(chan int, 1000)
	lock_map = make(map[int]int, MAXNUM)
	rwlock_map = make(map[int]int, MAXNUM)
	count1 := 0
	count2 := 0
	init_map(lock_map, rwlock_map)
	time1 := time.Now()
	for i := 0; i < LOCKNUM; i++ {
		go test1(lock_ch, i, lock_map, &lock_w)
	}
	go func() {
		lock_w.Wait()
		close(lock_ch)
	}()
	for i := range lock_ch {
		count1 += i
	}
	fmt.Printf("CHAN ID SUM %d\n", count1)

	time2 := time.Now()
	for i := 0; i < LOCKNUM; i++ {
		go test2(rwlock_ch, i, rwlock_map, &rwlock_w)
	}
	go func() {
		rwlock_w.Wait()
		close(rwlock_ch)
	}()
	for i := range rwlock_ch {
		count2 += i
	}
	fmt.Printf("CHAN ID SUM %d\n", count2)
	time3 := time.Now()
	fmt.Println("lock time:", time2.Sub(time1).String())
	fmt.Println("rwlock time:", time3.Sub(time2).String())
}

func init_map(a map[int]int, b map[int]int) { //初始化map
	for i := 0; i < MAXNUM; i++ {
		a[i] = i
		b[i] = i
	}
}

func test1(ch chan int, i int, mymap map[int]int, w *sync.WaitGroup) int {
	lock.Lock()
	defer lock.Unlock()
	ch <- i
	time.Sleep(time.Nanosecond)
	w.Done()
	return mymap[i%MAXNUM]
}

func test2(ch chan int, i int, mymap map[int]int, w *sync.WaitGroup) int {
	rwlock.RLock()
	defer rwlock.RUnlock()
	ch <- i
	time.Sleep(time.Nanosecond)
	w.Done()
	return mymap[i%MAXNUM]
}
