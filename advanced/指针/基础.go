package main

import "fmt"

/*
https://juejin.cn/post/7236213437801087031

什么时候在 Go 语言中使用指针？
1. 在函数之间传递大型数据结构
2. 修改函数内变量的值
3. 动态分配内存
*/

// 2. 修改函数内变量的值
func main2() {
	var x int = 10
	var p *int = &x

	fmt.Println("*p = ", *p) // *p =  10

	println(p, &x) // 0xc00007df08 0xc00007df08

	println(p == &x) // true

}

// 3. 动态分配内存
func main() {
	// 声明一个指向 int 类型的指针 p
	var p *int = new(int)
	*p = 100
	fmt.Println(*p) // 100

	var str *string
	fmt.Printf("str: %v\n", str) // 输出 p: <nil>

	str = new(string) // 分配一个string类型的内存，并将指针str指向该内存
	*str = "Hello, Go!"
	fmt.Printf("*str: %s\n", *str) // 输出 *str: Hello, Go!

	println("\n---------------------------------------------\n")
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中  等价于 b *int = &a
	fmt.Printf("type of b:%T\n", b)

	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	/*
		type of b:*int
		type of c:int
		value of c:10
	*/
}

// 1. 在函数之间传递大型数据结构
type worker struct {
	salary float64
	work   string
}

func changeWorker(p *worker) {
	p.salary = 999999.99
	p.work = "老板"
}

func main1() {
	// 引用
	w := &worker{
		salary: 1000.00,
		work:   "搬砖",
	}

	fmt.Println(w)      // &{1000 搬砖}
	fmt.Println(w.work) // 搬砖
	// 升职加薪
	changeWorker(w)
	fmt.Println(w) // &{999999.99 老板}

	// 不是引用
	w2 := worker{
		salary: 1000.00,
		work:   "搬砖2",
	}
	fmt.Println(w2) // {1000 搬砖2}

	// 引用
	changeWorker(&w2)
	fmt.Println(w2) //{999999.99 老板}

}
