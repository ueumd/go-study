package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var i int = 1
	var j byte = 2

	j = byte(i)

	fmt.Printf("j: %v\n", j)

	//var i int = 1
	//var j byte = 2
	//j = byte(i)
	//fmt.Printf("j: %v\n", j)

	s := fmt.Sprintf("是字符串:%s", "string")
	fmt.Println(s)

	var a byte = 'A'
	var b rune = 'A'
	var c rune = '中'
	fmt.Println(a) // 65
	fmt.Println(b) // 65
	fmt.Println(c) // 20013

	var aa rune
	fmt.Printf("%v, type: %T, char: %c", aa, aa, aa)

	fmt.Println()

	str1 := "hello"
	str2 := []byte{104, 101, 108, 108, 111}

	fmt.Println(str1)        // hello
	fmt.Printf("%s\n", str2) // hello
	fmt.Println(str2)        // [104 101 108 108 111]

	var num byte = 255
	fmt.Println(num)

	fmt.Println("------------\n")
	test()
	test2()

	fmt.Println("------------\n")
	test4()

	fmt.Println("------------\n")
	test5()
	fmt.Println("------------\n")
	test6()

	fmt.Println("------------\n")
	test7()

}

func test() {
	var a byte = 65
	// 8进制写法: var c byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var c byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 66
	fmt.Printf("a 的值: %c , b 的值: %c \n", a, b) // a 的值: A , b 的值: B
	// 或者使用 string 函数
	fmt.Println("a 的值:", string(a), ", b 的值:", string(b)) // a 的值: A , b 的值: B

	// 等价如下写法
	var A byte = 'A'
	var B uint8 = 'B'
	fmt.Printf("A 的值: %c , B 的值: %c", A, B) // A 的值: A , B 的值: B
}

func test2() {
	var a byte = 'A'
	var b rune = 'B'

	fmt.Printf("a 占用 %d 个字节， b 占用 %d 个字节", unsafe.Sizeof(a), unsafe.Sizeof(b))
	// a 占用 1 个字节， b 占用 4 个字节
}

func test3() {
	var str1 string = "hello"
	var str2 []byte = []byte{104, 101, 108, 108, 111}

	fmt.Println(str1)        // hello
	fmt.Printf("%s\n", str2) // hello
	fmt.Println(str2)        // [104 101 108 108 111]
}

func test4() {
	var str1 string = "hello"
	var str2 string = "你好"

	// golang的字符串都是以UTF-8格式保存，每个中文占用3个字节，因此使用 len() 获得两个中文文字对应的 6 个字节。
	fmt.Println(len(str1)) // 5
	fmt.Println(len(str2)) // 6

	// 字符个数
	fmt.Println(utf8.RuneCountInString(str1)) // 5
	fmt.Println(utf8.RuneCountInString(str2)) // 2

}

func test5() {
	char := '你'

	v1 := rune(char)
	v2 := byte(char)

	s1 := strconv.FormatInt(int64(v1), 2)
	s2 := strconv.FormatInt(int64(v2), 2)

	fmt.Printf("v1: %c, type: %T, %v", v1, v1, s1) // v1: 你, type: int32, 100111101100000
	fmt.Println()
	fmt.Printf("v2: %c, type: %T, %v", v2, v2, s2) // v2: `, type: uint8, 1100000

	/**
	v1: 你, type: int32, 100111101100000
	v2: `, type: uint8, 1100000
	*/
}

func test6() {
	var aa byte = 'A'
	fmt.Printf("aa: %v type: %T \n", aa, aa)

	var aaa string = "A"
	bbb := []byte(aaa)
	fmt.Printf("bbb: %v type: %T \n", bbb, bbb)

	fmt.Println("------------test6-----------")
	str := "你好"
	b := []byte(str)
	c := []rune(str)
	fmt.Printf("b: %v b: %c type: %T \n", b, b, b) // b: [228 189 160 229 165 189] b: [ä ½  å ¥ ½] type: []uint8

	fmt.Printf("c: %v c: %c type: %T \n", c, c, c) // c: [20320 22909] c: [你 好] type: []int32

	fmt.Println()

	str2 := []byte{228, 189, 160, 229, 165, 189, 239, 188, 140, 104, 101, 108, 108, 111}

	fmt.Printf("%s\n", str2) // hello
}

func test7() {
	str := "abcdef"

	//string 转[]byte
	b := []byte(str)

	fmt.Printf("b: %c \n", b) // b: [a b c d e f]

	//[]byte转string
	str = string(b)
	fmt.Printf("str: %s \n", str) // str: abcdef

	//string 转 rune
	r := []rune(str)
	fmt.Printf("r: %c\n", r) //r: [a b c d e f]

	//rune 转 string
	str = string(r)
	fmt.Printf("str: %s \n", str) // str: abcdef
}
