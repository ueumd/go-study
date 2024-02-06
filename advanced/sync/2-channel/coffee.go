package main

import (
	"fmt"
	"time"
)

/**
在goroutine coffee.go DEMO中使用咖啡机时A和B并发执行，会有交叉使用，而且会有插队的情况。

那如果现在还是想A人员优先使用，使用完后再告诉B队人员可以冲咖啡了，怎么做呢？

使用 channel 实现同步
*/

var makedCoffee = make(chan string)

var (
	ListA = []string{"小明", "小红", "小王", "宝宝", "娜娜", "拉拉", "琪琪"}
	ListB = []string{"威廉", "巴克", "查理", "哈利", "哈瑞", "米菲", "珍妮", "米奇"}
)

// 串行
func makeCoffee(list []string, from byte) {
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

func taskA() {
	makeCoffee(ListA, 'A')
	makedCoffee <- "A队已经冲好咖啡，B队你们可以冲咖了"
}

func taskB() {
	msg := <-makedCoffee
	fmt.Println(msg)

	makeCoffee(ListB, 'B')

	// 关闭通道
	close(makedCoffee)

	msg, ok := <-makedCoffee
	if ok == false {
		fmt.Println("通道已关闭")
	}

	fmt.Println(msg)
}

func main() {
	go taskA()
	go taskB()
	for {
	}
}
