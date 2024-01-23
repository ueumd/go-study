package main

import "fmt"

func main()  {

	// 1. 给一个类型定义别名，这种别名实际上为了代码的可读性强，本质上扔然是原始定义的类型

	type myByte = byte // 本质是 uint8
	var b myByte

	fmt.Printf("%T", b) // uint8

	// 2. 基于一个已有类型定义一个新类型

	// 注意和上面区别 没有 等号
	type myInt int
	var i myInt

	fmt.Printf("%T", i) // uint8main.myInt


	// 3. 定义结构体
	type courseTest struct {

	}

	// 4. 定义接口
	type callable interface {

	}

	// 5. 定义函数别名
	type handle func(str string)

}
