package main

import (
	"fmt"
	"sync"
)

/**
先进先出
*/

type Queue struct {
	list []interface{}
	size int
	lock sync.Mutex // 并发安全
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

// 入队
func (queue *Queue) Add(element interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面 复杂度O(n)
	queue.list = append(queue.list, element)

	queue.size += 1
}

// 出队
func (queue *Queue) Remove() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		return -1
	}

	// 先进先出
	value := queue.list[0]

	//切片内容
	newList := make([]interface{}, queue.size-1, queue.size-1)

	for i := 1; i < queue.size; i++ {
		// 老数组
		newList[i-1] = queue.list[i]
	}

	queue.list = newList

	queue.size -= 1
	return value
}

func main() {
	queue := new(Queue)

	queue.Add("Vue")
	queue.Add("React")
	queue.Add("Angular")
	queue.Add("jQuery")

	fmt.Println(queue.list)
	fmt.Println(queue.size)

	queue.Remove()

	fmt.Println(queue.list)
	fmt.Println(queue.size)

}
