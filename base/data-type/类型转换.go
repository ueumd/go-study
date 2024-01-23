package main

import "fmt"

func main() {
	/* ----------------- 1、整型和整型之间的转换  ----------------- */
	var a int8 = 20
	var b int16 = 40
	fmt.Println(int16(a) + b) // 60

	//2、浮点型和浮点型之间的转换
	var a1 float32 = 20
	var b1 float64 = 40
	fmt.Println(float64(a1) + b1) // 60

	//3、整型和浮点型之间的转换
	var a2 float32 = 20.23
	var b2 int = 40
	fmt.Println(a2 + float32(b2)) // 60.23

	//注意: 转换的时候建议从 低位 转换成 高位，高位 转换成 低位 的时候如果转换不成功就会溢出，如果和我们想的不一样

	var a3 int8 = 20
	var b3 int16 = 140
	fmt.Println(int16(a3) + b3) // 160

	fmt.Println(a3 + int8(b3)) // -96 溢出，如果和我们想的不一样

	/* ----------------- 1、整型和整型之间的转换  ----------------- */
	/**
	%表示格式化字符串输出
	printf("%c",a)；输出单个字符。
	printf("%d",a)；输出十进制整数。
	printf("%f",a)；输出十进制浮点数.
	printf("%o",a)；输出八进制数。
	printf("%s",a)；输出字符串。
	printf("%u",a)；输出无符号十进制数。
	printf("%x",a)；输出十六进制数。

	*/

	var i int = 20
	var f float64 = 12.456
	var t bool = true
	var bt byte = 'a'

	var strs = ""

	strs = fmt.Sprintf("%d", i)
	fmt.Printf("str type %T ,strs=%v  \n", strs, strs)

	strs = fmt.Sprintf("%f", f)
	fmt.Printf("str type %T ,strs=%v  \n", strs, strs)

	strs = fmt.Sprintf("%t", t)
	fmt.Printf("str type %T ,strs=%v  \n", strs, strs)

	strs = fmt.Sprintf("%c", bt)
	fmt.Printf("str type %T ,strs=%v \n", strs, strs)

}
