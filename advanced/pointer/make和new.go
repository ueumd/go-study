package main

import "fmt"

/*
在 Go 语言中
1. 对于值类型的声明，无需手动分配内存空间，因为它们在声明时已经默认分配了内存空间。
2. 对于引用类型的变量，在使用之前需要先进行声明，并为其分配内存空间，否则无法存储值。

- new 主要用于分配值类型的内存空间，
- make 主要用于分配引用类型（如 slice、map 和 channel）的内存空间
*/
func test_pt1() {
	var a *int

	fmt.Printf("a: %v \n", a) // a: <nil>

	// a 是一个空指针，野指针不能直接赋值
	// *a = 100
	// panic: runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(*a)

	a = new(int) // 开辟空间地址

	fmt.Println(*a) // 0

	*a = 100
	fmt.Println(*a) // 100
}

// go 中没有枚举类型
const (
	OK                  = 0
	FAIL                = 500
	ERROR_TAG_EXIST     = 4001
	ERROR_TAG_NOT_EXIST = 4002
	ERROR_TAG_ART_EXIST = 4003
)

func test_pt2() {
	var result map[int]string
	result = make(map[int]string)

	result[401] = "未受权"
	result[404] = "资源不存在"

	result[OK] = "OK"
	result[FAIL] = "FAIL"

	fmt.Println(result) //map[0:OK 401:未受权 404:资源不存在 500:FAIL]

	for k, v := range result {
		fmt.Printf("%d: %s \n", k, v)
	}

}

func main() {
	test_pt1()
	test_pt2()
}
