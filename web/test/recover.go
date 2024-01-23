package main

import "fmt"

func test_recover() {
	defer func() {
		fmt.Println("defer func")

		/**
		Go 语言还提供了 recover 函数，可以避免因为 panic 发生而导致整个程序终止
		recover 函数只在 defer 中生效
		*/
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}

func main() {
	// recover 捕获了 panic，程序正常结束
	// test_recover() 中的 after panic 没有打印，这是正确的，当 panic 被触发时，控制权就被交给了 defer
	
	test_recover()
	fmt.Println("after recover")

	/**
	defer func
	recover success
	after recover
	*/
}
