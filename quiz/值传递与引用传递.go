package main

import "fmt"

/**
https://blog.csdn.net/baolingye/article/details/111142386
*/

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice) // [1 2 3 4 5]

	// 函数内部扩容，外部切片会改变
	exchangeSlice(slice)
	fmt.Println(slice) // [2 4 6 8 10]

	// 函数内部进行了扩容
	exchangeSlice2(slice)
	fmt.Println(slice) // [2 4 6 8 10]

}

func exchangeSlice(slice []int) {
	for k, v := range slice {
		slice[k] = v * 2
	}
}

/*
*
在内部函数内如果发生了底层数组的扩容，
那么函数内的slice则会申请新的数组，在扩容后的slice做的变更也将不会再反馈到外部main函数的slice中
*/
func exchangeSlice2(slice []int) {
	slice = append(slice, 4, 5, 6) // 原始数组大小为4，超过4会扩容
	for k, v := range slice {
		slice[k] = v * 2
	}
	fmt.Println(slice) // [4 8 12 16 20 8 10 12]

}
