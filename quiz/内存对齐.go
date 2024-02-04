package main

import (
	"fmt"
	"unsafe"
)

type BadSt struct {
	A int32
	B int64
	C bool
}

type GoodSt struct {
	A int32
	C bool // 一个顺序调整就节省了8个字节
	B int64
}

func main() {
	bad := BadSt{A: 10, B: 20, C: false}
	fmt.Println(unsafe.Sizeof(bad)) //输出结果：24

	good := GoodSt{A: 10, B: 20, C: false}
	fmt.Println(unsafe.Sizeof(good)) //输出结果：16
}
