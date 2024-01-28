package main

import (
	"fmt"
	"image"
)

/**

https://docs.golangjob.cn/advanced-go-programming-book/ch1-basic/ch1-03-array-string-and-slice.html

数组有如下特点：

1. 数组是值类型，默认情况下作为参数传递给函数时会拷贝，不影响原数组。
2. 声明时，默认值为零值（0、""、false等）
3. 声明固定长度，不能动态变化
4. 元素的数据类型要相同
5. 占用内存连续

// 数组的创建有下面三种方式
[length]Type
[N]Type{value1, value2, ..., valueN}
[...]Type{value1, value2, ..., valueN}

因为数组的长度是数组类型的一个部分，不同长度或不同类型的数据组成的数组都是不同的类型，因此在Go语言中很少直接使用数组（不同长度的数组因为类型不同无法直接赋值）。
*/

func testArr() {
	// [length]Type
	var a [5]int   //  // 声明一个长度为5的int类型数组
	fmt.Println(a) // [0 0 0 0 0]

	// [N]Type{value1, value2, ..., valueN}
	b := [5]string{"Go", "C", "C++", "Vue", "React"}
	fmt.Println(b) // [    ]

	var c [5]bool
	fmt.Println(c) // [false false false false false]

	// 添加元素
	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println("a: ", a) // a:  [1 2 3 0 0]

	b = [5]string{"hello", "go", "vue"} // 前三个元素初始化为"hello", "go", "vue"
	fmt.Println("b: ", b)               // b:  [hello go vue  ]

	c = [5]bool{true, true}
	fmt.Println("c: ", c) // c:  [true true false false false]

	// 数组遍历
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d;  ", i, i) // a[0] = 0;  a[1] = 1;  a[2] = 2;  a[3] = 3;  a[4] = 4;
	}

	fmt.Println("\n")

	// 数组访问
	fmt.Println("a[1] = ", a[1]) // a[1] =  2

	fmt.Println("a.length", len(a), "b.length", len(b)) // a.length 5 b.length 5

	// 数组的长度根据初始化元素的数目自动计算
	// 当遇到[...]的声明时，其长度会被标记为nil，将在后续阶段进行自动推断。
	// [...]Type{value1, value2, ..., valueN}
	var x = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var y = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	fmt.Println("--------------------结构体数组--------------------")
	var line1 = [...]image.Point{{0, 0}, {1, 1}}
	var line2 = [...]image.Point{image.Point{X: 100, Y: 100}, image.Point{X: 200, Y: 200}}
	fmt.Println(line1, line2)

}

func testArr2() {
	var d [0]int       // 定义一个长度为0的数组
	var e = [0]int{}   // 定义一个长度为0的数组
	var f = [...]int{} // 定义一个长度为0的数组
	fmt.Println(d, e, f)

	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}

	fmt.Println(unknown2, unknown1)

	// 管道数组
	var chanList = [2]chan int{}
	fmt.Println(chanList)
}

func testArrStr() {
	// 字符数组
	var data = [...]byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'}
	fmt.Println(data) // [104 101 108 108 111 44 32 119 111 114 108 100]

	for i := 0; i < len(data); i++ {
		fmt.Printf("%s, ", string(data[i])) // h, e, l, l, o, ,,  , w, o, r, l, d,
	}
}

// 数组是值传递
// 值传递, 复制数组
func changeArr(arr [4]int) {
	for key, value := range arr {
		arr[key] = value + 1
	}
}

// 为了避免复制数组，一般会传递指向数组的指针
// 引用传递
func changeArr2(arr *[4]int) {
	for key, value := range arr {
		arr[key] = value + 1
	}
}

func main() {
	// testArr()
	testArrStr()

	println("------------")

	arr := [...]int{100, 200, 300, 400}
	fmt.Println("arr: ", arr)

	// 值传递，不会改变原始数组
	changeArr(arr)
	fmt.Println("arr:", arr) // arr: [100 200 300 400]

	// 引用传递，会改变原始数组
	changeArr2(&arr)
	fmt.Println("arr:", arr) // arr: [101 201 301 401]
}
