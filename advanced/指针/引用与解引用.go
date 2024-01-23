package main

import "fmt"

// 值拷贝
func double(x int) {
	x += x
}

/*
var x int = 10
var p *int = &x
访问存储在指针内存地址中的变量的值，请再次使用 * 符号
fmt.Println(*p) // prints 10
*/

// *int 是对 &int  的解引用

func double2(x *int) {
	*x += *x
}

func main() {
	var a = 3

	// 传的是值
	double(a)
	fmt.Println("a = ", a) // 3

	// 传引用
	double2(&a)
	fmt.Println("a = ", a) // a = 6

	p := &a
	double2(p)
	fmt.Println("a = ", a) // a =  12

	println("p == &a", p == &a) // true

	// 指针地址
	println("p: ", p)   // p:   0xc00007dec8
	println("&a: ", &a) // &a:  0xc00007dec8
}
