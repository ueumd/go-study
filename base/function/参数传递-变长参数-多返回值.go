package main

import (
	"errors"
	"fmt"
)

/*
https://cloud.tencent.com/developer/article/2232168
1. 按值传参
2. 引用传递
3. 可变参数
4. 多返回值
5. 命名返回值
6. 传入函数
*/

// ------------------- 1.按值传参 -------------------------------
func add(a, b int) int {
	a *= 2
	b *= 3
	return a + b
}

// ------------------- 2.引用传递 -------------------------------
func addRef(a, b *int) int {
	*a *= 2
	*b *= 3
	return *a + *b
}

// ------------------- 3.可变参数 -------------------------------
// ...type 格式的类型只能作为函数的参数类型存在，并且必须是函数的最后一个参数。
func myFunc(numbers ...int) {
	for _, number := range numbers {
		// fmt.Println(number)
		fmt.Printf("%d, ", number)
	}
	fmt.Println("\n")
}

// ...type 本质上是一个切片，也就是 []type，这也是为什么上面的参数 numbers 可以用 for 循环来获取每个传入的参数值。
// 假如没有 ...type 这样的语法糖，要实现同样的功能，开发者将不得不这么写：
func myFunc2(numbers []int) {
	for _, number := range numbers {
		fmt.Printf("%d, ", number)
	}
	fmt.Println("\n")
}

// ------------------- 4.多返回值 -------------------------------
func addResult(a, b *int) (int, error) {
	if *a < 0 || *b < 0 {
		err := errors.New("只支持非负整数相加")
		return 0, err
	}
	*a *= 2
	*b *= 3
	return *a + *b, nil
}

// ------------------- 5.命名返回值
func addResult2(a, b *int) (c int, err error) {
	if *a < 0 || *b < 0 {
		err := errors.New("只支持非负整数相加")
		return 0, err
	}
	*a *= 2
	*b *= 3
	c = *a + *b
	return
}

// ------------------- 5.命名返回值
func testFunc(fns ...func(num int)) {
	for _, fn := range fns {
		myNum := 100
		fn(myNum)
	}
}

func testFunc2(callback func(num int)) {
	callback(100)
}

func main() {
	fmt.Println("---------------传 传 递-----------------")
	x, y := 1, 2
	z := add(x, y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z) // add(1, 2) = 8

	fmt.Println("---------------引用传递-----------------")
	a, b := 1, 2
	c := addRef(&a, &b)
	fmt.Printf("addRef(%d, %d) = %d\n", a, b, c) // addRef(2, 6) = 8

	fmt.Println("---------------可变参数-----------------")
	myFunc(100, 200, 300, 400)

	slice := []int{1, 2, 3, 4, 5}
	myFunc(slice...)
	myFunc(slice[1:3]...)

	// ...type 本质上是一个切片，也就是 []type
	myFunc2(slice[1:3])

	fmt.Println("---------------多返回值-----------------")
	x1, y1 := 1, 2
	z1, err := addResult(&x1, &y1)
	if err == nil {
		fmt.Printf("addResult(%d, %d) = %d\n", x1, y1, z1) //addResult(2, 6) = 8
	}

	fmt.Println("---------------命名回值-----------------")
	z1, err = addResult2(&x1, &y1)
	if err == nil {
		fmt.Printf("addResult2(%d, %d) = %d\n", x1, y1, z1) //addResult2(4, 18) = 22
	}

	fmt.Println("---------------函数当参数-----------------")
	testFunc(func(num int) {
		fmt.Println(num)
	}, func(num int) {
		fmt.Println("2------", num)
	})

	testFunc2(func(num int) {
		fmt.Println("callback value: ", num)
	})
}
