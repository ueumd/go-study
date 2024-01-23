package main

import (
	"fmt"
	"reflect"
)

func add[T int | int32 | float64 | string](a, b T) T {
	return a + b
}

func init() {
	fmt.Println(add(1, 3))
	fmt.Println(add("A", "B"))

	var int32A int32 = 3
	var int32B int32 = 4

	res := add(int32A, int32B)
	fmt.Println(res)

	fmt.Println("\n")
}

// 获取指针
func ptr[T any](in T) *T {
	return &in
}

// 获取指针值
func ptrValue[T any](in *T) T {
	return *in
}

func main() {

	intPtr1 := ptr[int](100)
	fmt.Println(*intPtr1, reflect.TypeOf(intPtr1))

	intPtr2 := ptr(200)
	fmt.Println(*intPtr2, reflect.TypeOf(intPtr2))

	strPtr := ptr("abc")
	fmt.Println(reflect.TypeOf(strPtr))

	strValue := ptrValue(strPtr)

	fmt.Println(strValue, reflect.TypeOf(strValue))

}
