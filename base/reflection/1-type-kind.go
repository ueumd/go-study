package main

import (
	"fmt"
	"reflect"
)

/**
- 反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
- 反身会将匿名字段作为独立字段（匿名字段本质）
- 想要利用反射修改对象状态，前提是interface.data是settable,即pointer-interface
- 通过反射可以“动态”调用方法
*/

type UnKnownType struct {
	S1 string
	S2 string
	S3 string
}

func (n UnKnownType) string() string {
	return n.S1 + " & " + n.S2 + " & " + n.S3
}

var lang interface{} = UnKnownType{"Go", "C", "Python"}

func main() {
	langType := reflect.TypeOf(lang)
	fmt.Println(langType) //main.UnKnownType  反射的类型

	// Kind 种类（Kind）指的是对象归属的品种
	fmt.Println(langType.Name()) // UnKnownType
	fmt.Println(langType.Kind()) // struct

	langValue := reflect.ValueOf(lang)
	fmt.Println(langValue) // {Go C Python}

	// 字段
	for i := 0; i < langValue.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, langValue.Field(i))
		/**
		Field 0: Go
		Field 1: C
		Field 2: Python
		*/
	}

}
