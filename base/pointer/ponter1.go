package main

import "fmt"

func pp1() {
	var a = 10
	var b = &a

	fmt.Printf("a: %d,	ptr: %p \n", a, &a) // a:10 ptr:0xc0000100a8
	fmt.Printf("b: %v,	type: %T \n", b, b) // b:0xc0000100a8 type:*int
	fmt.Println("b的地址：", &b)               // 0xc00005e020

	/**
	a: 10,  ptr: 0xc00000a0b8
	b: 0xc00000a0b8,        type: *int
	b的地址： 0xc00005e020
	*/
}

/**
指针
空指针： 没有初始化的指针。var p *int
野指针： 被一片无效的地址空间初始化
*/

func p1() {
	var pp int // 变量pp 默认值是变量类型的初始值 int 即 0，会自动开辟空间
	fmt.Println(pp)

	// 空指针
	var p *int //声明了一个指针变量p 内存空间值为nil 空指针

	// 指针变量 p 没有开辟内存，会报错
	*p = 10 // 把nil的内存空间值变成 10，这就很奇怪！！！

	fmt.Println(*p) // invalid memory address or nil pointer dereference 无效的内存地址 或空指针引用
}

func p2() {
	var p int = 10
	fmt.Println("&p = ", &p) // &p =  0xc00000a098

	// 野指针
	//var p2 *int = 0xc00000a098
	//fmt.Println(*p)
}

func p3() {

	// heap上申请一片内存空间

	var p *string

	p = new(string) // 编译器申请内存空间，默认值为变量类型的初始值 nil

	*p = "Abcdef"

	fmt.Printf("%s\n", *p)
	fmt.Printf("%q\n", *p)
}

/*
*
new 和 make的区别
new 可以申请内存空间 make也可以
make 通常用在 slice map

区别：
new 返回的是这个值的地址 指针
make 返回的是指定类型的实例
*/
func main() {
	var p *int = new(int)
	*p = 10

	var msg map[string]string
	if msg == nil {
		fmt.Println("map的默认值为nil")
	}

	var err error
	if err == nil {
		fmt.Println("error的默认值为nil")
	}

	var slice []string
	if slice == nil {
		fmt.Println("切片的默认值为nil")
	}

	// go 语言中 nil是唯一可以用来表示部分类型的零值的 标识符，它可以代表许多不同内存布局的值

	var info map[string]string = make(map[string]string)
	if info == nil {
		fmt.Println("map的默认值为nil")
	}
	info["c"] = "bobby"
}
