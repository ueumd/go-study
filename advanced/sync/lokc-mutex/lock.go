package main

import "sync"

import (
	"fmt"
)

/**
Go 标准库中提供了 sync.RWMutex 互斥锁类型及其四个方法：

Lock 加写锁
Unlock 释放写锁
RLock 加读锁
RUnlock 释放读锁
*/

/*
*
struct field使用sync.Mutex 跟 *sync.Mutex的区别
*/
type Person struct {
	// Person进行初始化之后,Mu也会初始一个实例, Mu的lock就会初始化为默认0, 开锁状态
	Mu  sync.Mutex
	Age int
}

type Person2 struct {
	// 使用指针类型, 初始化时Mu只会初始化个nil 出来,在nil调用Lock(),runtime会panic出内存地址不合法的错误
	// 必须手动给Mu初始化一个实例 xx.Mu = &sync.Mutex{}
	Mu  *sync.Mutex
	Age int
}

func (p2 *Person2) change() {
	p2.Mu.Lock()
	defer p2.Mu.Unlock()
	p2.Age++
}

func main() {
	p2 := Person2{}
	p2.Mu = &sync.Mutex{}

	p2.change()
	p2.change()
	p2.change()
	p2.change()
	fmt.Println(p2.Age) // 4
}

func main111() {

	/*
		如果是使用普通值类型呢,使用,我们使用Person{}或者
		new()函数对Person进行初始化之后,Mu也会初始一个实例,
		Mu的lock就会初始化为默认0, 开锁状态
	*/
	me := Person{}
	me.Mu.Lock()
	me.Age++
	me.Mu.Unlock()
	fmt.Println(me.Age)

	/*
		如果是使用指针类型的话呢, 初始化时Mu只会初始化个nil
		出来,在nil调用Lock(),runtime会panic出内存地址不合法的错误
	*/
	me2 := Person2{}
	// me2.Mu = &sync.Mutex{} //我们手动初始化一个实例赋值给Mu,就不会报错了
	me2.Mu.Lock()
	me2.Age++
	me2.Mu.Unlock()
	fmt.Println(me2.Age)

	/*
		无论是值类型还是指针类型,它们调用的Lock() Unlock()都是一样的:
		func (m *Mutex) Lock()
		func (m *Mutex) Unlock()
	*/
}
