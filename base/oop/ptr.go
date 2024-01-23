package main

import "fmt"

type student struct {
	age  int
	name string
}

/**
go方法传参时指针类型和非指针类型的区别
*/

// 修改不了实参的成员属性 对传入数据进行拷贝 对stu结构体进行拷贝，不会影响原来的的结体构
func changeName(stu student) {
	stu.name = "Haha"
}

// 指针地址 引用
func changeName2(stu *student) {
	stu.name = "OOO"
}

func main() {
	stu := student{
		name: "Tom",
		age:  18,
	}

	changeName(stu)

	fmt.Println(stu.name) // tom

	// 传入指针地址
	changeName2(&stu)
	fmt.Println(stu.name) // OOO
}
