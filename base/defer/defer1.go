package main

import "fmt"

// 匿名返回 5
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 命名返回 6
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 命名返回值，但是返回的变量名 和 命名不一样，相当于 x 赋值给 y，所以返回5
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		// 这里的x是参数传进来的x，如 y , 再y++
		x++
	}(x)

	return 5
}

func f42() (x int) {
	defer func(y int) {
		fmt.Println("y = ", y)
		y++
	}(x)

	fmt.Println("x = ", x)
	return 5
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5

	fmt.Println(f42())
	/**
	x =  0
	y =  0
	5
	*/
}
