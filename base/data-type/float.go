package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	var num float64 = 3.1415925535
	// %v 原样输出 %f 输出float类型  %.2f 保留2位小数,四舍五入
	fmt.Printf("%v -- %f -- %.2f \n", num, num, num)
	// 3.1415925535 -- 3.141593 -- 3.14

	fmt.Printf("%v -- %f -- %.4f \n", num, num, num)
	// 3.1415925535 -- 3.141593 -- 3.1416

	// 64位系统中 浮点数默认是64位  %T 打印数据类型

	f1 := 3.13456456
	fmt.Printf("%f--%T \n", f1, f1) // 3.134565--float64

	//************* golang 科学计数法 *************//

	// f2 = 3.14 * 10^2 = 3.14 * 100
	var f2 = 3.14e2
	fmt.Printf("%f--%T \n", f2, f2) // 314.000000--float64
	fmt.Printf("%v--%T \n", f2, f2) // 314--float64

	// f3 = 3.14 / 10^2 = 3.14 / 100
	var f3 = 3.14e-2
	fmt.Printf("%f--%T \n", f3, f3) // 0.031400--float64
	fmt.Printf("%v--%T \n", f3, f3) // 0.0314--float64

	//************* float 精度丢失问题 *************//
	d := 1129.6
	fmt.Println(d * 100) //112959.99999999999

	var m1, m2 = 8.2, 3.8
	fmt.Println(m1 - m2) // 期望是4.4 输出 4.3999999999999995

	fmt.Println(decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(m2))) // 4.4

	// 解决 使用第三方包
	var num1, num2 = 3.1, 4.2
	fmt.Println(num1 + num2) // 7.300000000000001

	num3 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Println(num3) // 7.3

	// int 类型转 float类型
	a := 10
	b := float64(a)
	fmt.Printf("a 的类型是 %T, b的类型是 %T \n", a, b) // a 的类型是 int, b的类型是 float64

	// float类型 转 int （会丢失，不建议）

}
