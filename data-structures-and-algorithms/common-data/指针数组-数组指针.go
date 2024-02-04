package main

import "fmt"

/**
https://zhuanlan.zhihu.com/p/137443361

数组指针 存放的是数组的指针地址
var arrPtr *[size] Type

指针数组 所有的元素是指针地址
var ptrArr [size] *Type
*/

/*
*
数组指针 存放的是数组的指针地址
var arrPtr *[size] Type
*/
func arrPointer() {

	// 数组指针
	var arrPtr *[4]int
	var list = [4]int{1, 2, 3, 4}
	arrPtr = &list
	fmt.Println("arrPtr: ", arrPtr) // arrPtr:  &[1 2 3 4]

	fmt.Println(arrPtr == &list)               // true
	fmt.Printf("arrPtr address: %p\n", arrPtr) // arrPtr address: 0xc000012220
	fmt.Printf("list   address: %p\n", &list)  // list   address: 0xc000012220

	// 通过指针访问数组的第一个元素
	/**
	Golang 中 * 寻址运算符和 [] 中括号运算符的优先级是不同的！
	[] 中括号是初等运算符
	寻址运算符是单目运算符

	*/
	// fmt.Println(*arrPtr[1]) 报错

	// 不过因为 * 在Golang中，建立了 arrPtr := &list 这种类似地址关系后，* 允许不写。
	fmt.Println((*arrPtr)[1]) // 2

	// 可以直接写成 arrPtr[1]
	fmt.Println(arrPtr[1]) // 2
}

/*
*
指针数组 所有的元素是指针地址
var ptrArr [size] *Type
*/
func pointerArr() {
	var ptrArr [4]*int
	a, b, c, d := 1, 2, 3, 4

	// int，float，bool，string，array，struct都属于值类型，数据拷贝时都是值拷贝，拷贝副本过来
	list := [4]int{a, b, c, d}

	ptrArr = [4]*int{&a, &b, &c, &d}

	fmt.Println("list: ", list) // list:  [1 2 3 4]

	// 存的都是元素的地址
	fmt.Println("ptrArr: ", ptrArr) // ptrArr:  [0xc00000a128 0xc00000a130 0xc00000a138 0xc00000a140]

	list[0] = 100
	fmt.Println("list: ", list) // list:  [100 2 3 4]
	fmt.Println("a: ", a)       // a:  1

	fmt.Println("*prtArr[0]: ", *ptrArr[0]) // *prtArr[0]:  1

	*ptrArr[0] = 10000
	fmt.Println("a: ", a) // a:  10000
	// ptrArr[0] 存的是a的地址，所以会跟着变
	fmt.Println("*prtArr[0]: ", *ptrArr[0]) // *prtArr[0]:  10000

}

func main() {
	arrPointer()
	fmt.Println("----------------------------------------\n")
	pointerArr()
}
