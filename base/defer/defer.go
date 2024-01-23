package main

import "fmt"

func defer1()  {
	fmt.Println("test 1111")

	defer fmt.Println("test defer 1111")
	defer fmt.Println("test defer 2222")

	fmt.Println("test 2222")

	/**
	test 1111
	test 2222
	test defer 2222
	test defer 1111
	*/
}

func defer2()  {
	x := 10

	// 压栈 压栈时参数进行了值拷贝，不受x++影响
	defer func(a int) {
		fmt.Println(a) // 10
	}(x)

	x++ // x 压栈
}

func defer3()  {
	x := 10

	// 进行了压栈，但并不是传参
	defer func() {
		// 闭包的概念
		fmt.Println(x) // 11
	}()

	x++
}

func defer4()  {
	x := 10

	// 进行了压栈，引用传递 指针地址
	defer func(x *int) {
		// 闭包的概念
		fmt.Println(*x) // 11
	}(&x)

	x++
}


// 函数返回过程
func deferFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i

	/**
	延迟函数的执行正是在return之前，即加入defer后的执行过程如下：
	result = i
	result++
	return
	所以上面函数实际返回i++值 为2
	 */
}

// 主函数拥有匿名返回值，返回字面量
func foo() int {
	var i int

	defer func() {
		i++
	}()

	return 1

	/**
	返回一个局部变量，同时defer函数也会操作这个局部变量。
	对于匿名返回值来说，可以假定仍然有一个变量存储返回值，假定返回值变量为“anony”，上面的返回语句可以拆分成一下过程：
	anony=i
	i++
	return

	由于i是整形，会将值拷贝给anony，所以defer语句中修改i值，对函数返回值不造成影响，最终返回1

	 */
}

// 主函数拥有具名返回值
/**
主函数声明语句中带名字的返回值，会被初始化成一个局部变量，函数内部可以像使用局部变量一样使用该返回值。如果defer语句操作该返回值，可能会改变返回结果。
 */
func foo2() (ret int) {
	defer func() {
		ret++
	}()

	return 0

	/**
	上面的函数拆解出来，如下所示：
	ret = 0
	ret++
	return

	函数真正返回前，在defer中对返回值做了+1操作，所以函数最终返回1
	*/
}
func main1() {
	// defer1()
	defer2()
	defer3();
	defer4();

	fmt.Println(deferFuncReturn()) // 2
	fmt.Println(foo()) 			   // 1
	fmt.Println(foo2())			   // 1
}

func main()  {
	fmt.Println("test 1111")

	defer fmt.Println("test defer 1111")
	defer fmt.Println("test defer 2222")

	panic("")

	fmt.Println("test 2222")

	/**
	test 1111
	test defer 2222
	test defer 1111
	panic:

	goroutine 1 [running]:
	main.main()
			go-test/defer.go:135 +0xf8

	Process finished with the exit code 2
	 */

}