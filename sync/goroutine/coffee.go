package main

import (
	"fmt"
	"time"
)

var (
	ListA = []string{"小明", "小红", "小王", "宝宝", "娜娜", "拉拉", "琪琪"}
	ListB = []string{"威廉", "巴克", "查理", "哈利", "哈瑞", "米菲", "珍妮", "米奇"}
)

// 串行
func task(list []string, from byte) {
	i := 0
	for {
		if i >= len(list) {
			if from == 'A' {
				fmt.Println("A队人员已排队完使用咖啡机 \n")
			} else if from == 'B' {
				fmt.Println("B队人员已排队完使用咖啡机~~~")
			}

			break
		}

		if from == 'A' {
			fmt.Printf("A队[%d]: %s 使用咖啡机 \n", i+1, list[i])
		} else if from == 'B' {
			fmt.Printf("B队[%d]: %s 使用咖啡机 \n", i+1, list[i])
		}

		i++
		time.Sleep(time.Second)
	}
}

func taskList(i int) {
	fmt.Printf("A队[%d]: %s 使用咖啡机 \n", i, ListA[i])
	time.Sleep(time.Second)
}

func main1() {
	// 同时启动7个子协程. 并发执行，执行先后顺序无法保证
	for i := 0; i < len(ListA); i++ {
		go taskList(i)
	}
	for {
	}
}

func main() {
	// 串行
	//task(ListA, 'A')
	//task(ListB, 'B')

	// 并发使用一台咖啡机
	go task(ListA, 'A')
	go task(ListB, 'B')
	for {

	}
}
