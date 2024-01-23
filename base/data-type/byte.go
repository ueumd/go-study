package main

import "fmt"

func main() {

	// 字符属于 int 类型

	var a = 'a'
	// ASCII 编码值
	//%v 按值的本来值输出
	fmt.Printf("a: %v, type: %T \n", a, a) // a: 97, type: int32

	// 原样输出字符串  %c 输出单个字符
	fmt.Printf("a: %c, type: %T \n", a, a) // a: a, type: int32

	// 定义字符串输出字符
	var str = "this"
	fmt.Printf("value: %v, 原样输出：%c,  type: %T \n", str[2], str[2], str[2]) // value: 105, 原样输出：i,  type: uint8

	// 一个汉字占用3个字节（utf-8），一个字母占用1个字节
	// golang中汉字属于utf-8编码，占3个字节，如果是gbk编码占2个字节

	// unsafe.Sizeof() 没法查看string类型数据所占用的空间

	var str2 = "你好go"

	fmt.Println(len(str2)) // 8= 3+3+1+1

	/*******     汉字    ******/
	// 定义一个汉字字符
	// http://mytju.com/classcode/tools/encode_utf8.asp
	// golang汉字使用utf8编码，编码的值就是int 类型
	var c1 = '国'
	fmt.Printf("c: %c, v: %v,  type: %T \n", c1, c1, c1) // c: 国, v: 22269,  type: int32

	// v: 22269 Unicode编码10进制

	/**
	当需要处理中文、日文或者其他复合字符时，则需要用到|rune 类型。rune 类型实际是一个int32。

	Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
	*/

	var str3 = "golang"
	// var str3 = "你好golang"
	for i := 0; i < len(str3); i++ {
		// 如果有汉字就不能正常输出
		fmt.Printf("%v(%c)", str3[i], str3[i]) // 103(g)111(o)108(l)97(a)110(n)103(g)
	}

	fmt.Println("\n")

	var str4 = "你好golang"
	for _, val := range str4 {
		fmt.Printf("%v(%c)", val, val) // 20320(你)22909(好)103(g)111(o)108(l)97(a)110(n)103(g)
	}

	fmt.Println("\n")

	/***** 修改字符串 ****/
	/**
	要修改字符串，需要先将其转换成[]rune 或[]byte。，完成后再转换为 string。无论哪种转换,都会重新分配内存，并复制字节数组。
	*/

	// 字母
	var s1 = "big"
	// 强制
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1)) // pig

	// 包含汉字
	s2 := "白萝卜"
	runeS2 := []rune(s2)

	runeS2[0] = '红'
	fmt.Println(string(runeS2)) // 红萝卜

}
