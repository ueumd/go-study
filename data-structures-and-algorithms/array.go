package main

import (
	"fmt"
	"sync"
)

/*
可变长数组
*/

type Array struct {
	cache []interface{}
	len   int
	cap   int
	lock  sync.Mutex
}

/*
*
初始化数组
长度为 len
容量为 cap
*/
func Make(len, cap int) *Array {

	// 给array开辟空间
	array := new(Array)

	if len > cap {
		panic("len large than cap")
	}

	// 元数据
	list := make([]interface{}, cap, cap)
	array.cache = list

	array.cap = cap
	array.len = 0
	return array
}

func (self *Array) Len() int {
	return self.len
}

func (self *Array) Cap() int {
	return self.cap
}

// 添加元还给
func (array *Array) Append(element interface{}) {
	array.lock.Lock()
	defer array.lock.Unlock()

	// 长度等于容量，需要扩容
	if array.len == array.cap {

		// 容量不够，扩容为长度的2倍
		newCap := 2 * array.len

		// 如果之前的容量为0，那么新容量为1
		if array.cap == 0 {
			newCap = 1
		}

		newArray := make([]interface{}, newCap, newCap)

		// 老数组元素拷贝到新数组元素里
		for key, val := range array.cache {
			newArray[key] = val
		}

		// 替换数组
		array.cache = newArray
		array.cap = newCap
	}

	// 添加元素
	array.cache[array.len] = element

	// 真实长度加1
	array.len += 1
}

/*
*
添加多个元素
*/
func (array *Array) AppendMany(element ...interface{}) {
	for _, value := range element {
		array.Append(value)
	}
}

/*
*
获取指定下标元素
*/
func (array *Array) Get(index int) interface{} {

	// 越界处理
	if array.len == 0 || index >= array.len {
		panic("index over len")
	}

	return array.cache[index]
}

func main() {
	arr := Make(0, 3)
	fmt.Println("cap: ", arr.Cap(), " len: ", arr.Len(), " array: ", arr.cache)

	arr.Append(1)
	arr.Append('A')
	arr.Append('B')

	arr.AppendMany(3, "C")
	//arr.AppendMany([]int{1, 2, 3})

	fmt.Println("cap: ", arr.Cap(), " len: ", arr.Len(), " array: ", arr.cache)
}
