package main

import "fmt"

/*
*
数组的大小是固定的，当元特别多时，固定数组无法储存那么多值，所以可变长数组出现了
在Golang中可变长数组被内置在语言里面：切片slice

	type slice struct {
	    array unsafe.Pointer
	    len   int  // 切片长度
	    cap   int  // 切片容量
	}
*/

func main() {
	array := make([]int, 0, 2)

	fmt.Println("cap: ", cap(array), " len: ", len(array), " array: ", array) // cap:  2  len:  0  array:  []

	// 虽然 append 但是没有赋予原来的变量 array， 所以 以下依就是 cap:  2  len:  0  array:  []
	_ = append(array, 1)
	fmt.Println("cap: ", cap(array), " len: ", len(array), " array: ", array)

	_ = append(array, 1)
	fmt.Println("cap: ", cap(array), " len: ", len(array), " array: ", array)

	_ = append(array, 1)
	fmt.Println("cap: ", cap(array), " len: ", len(array), " array: ", array)

	fmt.Println("-----------------------------------------\n")

	// 赋予回原来的变量

	array = append(array, 1)
	fmt.Println("cap: ", cap(array), " len: ", len(array), " array: ", array) // cap:  2  len:  1  array:  [1]

	array = append(array, 2)
	fmt.Println("cap: ", cap(array), " len: ", len(array), " array: ", array)

	array = append(array, 3)
	fmt.Println("cap: ", cap(array), " len: ", len(array), " array: ", array)

	array = append(array, 4)
	fmt.Println("cap: ", cap(array), " len: ", len(array), " array: ", array)

	/**

	Golang的切片无法原地append，每次添加元素时返回新的引用地址，必须把该引用重新赋予之前的切片变量。
	并且，当容量不够时，会自动倍数递增扩容。事实上，Golang在切片长度大于1024后，会以接近于1.25倍进行容量扩容。

	具体可参考标准库runtime下的slice.go文件。

	cap:  2  len:  0  array:  []
	cap:  2  len:  0  array:  []
	cap:  2  len:  0  array:  []
	cap:  2  len:  0  array:  []
	-----------------------------------------

	cap:  2  len:  1  array:  [1]
	cap:  2  len:  2  array:  [1 2]
	cap:  4  len:  3  array:  [1 2 3]
	cap:  4  len:  4  array:  [1 2 3 4]

	*/

}
