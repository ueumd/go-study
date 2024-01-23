package main

import "fmt"

/**
冒泡排序算法是稳定的，因为如果两个相邻元素相等，是不会交换的，保证了稳定性的要求。


举个简单例子，冒泡排序一个 4 个元素的数列：4 2 9 1：

[]表示排好序 {}表示比较后交换的结果

i = 3; j = 0 & j <3; 第一轮： 4 2 9 1 从第一个数开始，4 比 2 大，交换 4，2			{2 4} 9 1
i = 3; j = 1 & j <3; 第一轮： {2 4} 9 1  接着 4 比 9 小，不交换			 			{2 4} 9 1
i = 3; j = 2 & j <3; 第一轮： 2 {4 9} 1  接着 9 比 1 大，交换 9，1  	 			2 4 {1 9}
第一轮结果： 2 4 1 [9]

第二轮开始：2 4 1 [9] 从第一个数开始，2 比 4 小，不交换
第二轮： {2 4} 1 [9] 接着 4 比 1 大，交换 4，1
第二轮结果： 2 1 [4 9]

第三轮开始：2 1 [4 9] 从第一个数开始，2 比 1 大，交换 2，1
第三轮结果： 1 [2 4 9]

结果： [1 2 4 9]

*/

func BubbleSort(arr []int) {
	n := len(arr)
	didSwap := false
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				didSwap = true
			}
		}

		if !didSwap {
			break
		}
	}
}

func BubbleSort2(arr []int) {
	n := len(arr)
	didSwap := false

	// 进行n-1轮迭代
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				didSwap = true
			}
		}

		// 如果在每一轮中没有交换过，说明不需要排序，直接结束
		if !didSwap {
			break
		}
	}
}

func main() {
	// list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	list := []int{4, 2, 9, 1}
	BubbleSort2(list)

	fmt.Println(list)
}
