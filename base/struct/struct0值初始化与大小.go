package main

import (
	"fmt"
	"unsafe"
)

type course struct {
	name  string
	price int
	url   string
}

func main() {
	var c1 course = course{}
	var c2 course
	var c3 *course = new(course)
	var c4 *course = &course{} // 等价于上面
	// var c5 *course 空指针 只声明指针变量c5，但是没有分配内存空间

	fmt.Println("零值初始化")
	fmt.Println(c1.price)
	fmt.Println(c2.price)
	fmt.Println(c3.price)
	fmt.Println(c4.price)

	// 结构体大小
	fmt.Println(unsafe.Sizeof(1)) // 8

	cc := course{"haha", 16, "golang.dev"}
	fmt.Println(unsafe.Sizeof(cc)) // 16 + 8 + 16 = 40
	fmt.Println(unsafe.Sizeof(c4)) // 8

	/**
	string 内部是一个结构
	*/
	fmt.Println(unsafe.Sizeof("abcdef haha hello")) // 16
}
