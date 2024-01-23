package main

import "fmt"

func main1() {
	fmt.Println("before panic")
	//
	//panic 会中止当前执行的程序，退出
	panic("crash")

	fmt.Println("after panic")
}

func main2() {
	arr := []int{1, 2, 3}
	// 越界
	fmt.Println(arr[4])

}

func main() {

	// panic 会导致程序被中止，但是在退出前，会先处理完当前协程上已经defer 的任务，执行完成后再退出。效果类似于 java 语言的 try...catch。
	defer func() {
		fmt.Println("defer func")
	}()

	arr := []int{1, 2, 3}
	// 越界
	fmt.Println(arr[4])

}
