package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

/**
Go语言命名规则：
大写字母开头的变量是可导出的，也就是其它包可以读取的，是公用变量；小写字母开头的就是不可导出的，是私有变量。
大写字母开头的函数也是一样，相当于JAVA中class里的带public关键词的公有函数； 小写字母开头的就相当于有private关键词的私有函数。
*/
// 大写为对外可以访问
func Test() {
	/*
	   default value
	   int    0
	   bool   false
	   string ''
	*/
	var power int
	power = 9000
	fmt.Printf("It's over %d\n", power)
}

/*
在Go中字符串是不可变的，例如下面的代码编译时会报错：cannot assign to s[0]
var s string = "hello"
s[0] = 'c'
相改变如下：
*/
func StringStudy() {
	s := "hello"
	c := []byte(s) //转为字符
	c[0] = 'c'
	s2 := string(c) // 再转换回 string 类型

	fmt.Println(s2)
}

func PrintStrLen() {
	str := "hello世界"
	// fmt.Println(len(str)) // 11 正确打印需要"unicode/utf8"包
	fmt.Println(utf8.RuneCountInString(str)) // 7

	err := errors.New("err: emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}

func definVar() {

	/**
	var a int         // a的值为: 0      类型为: int
	var b string      // b的值为: ""     类型为: string
	var c interface{} // c的值为: nil    类型为: interface{}
	var d sync.Pool   // d的值为: Pool{} 类型为: sync.Pool
	var e *sync.Pool  // e的值为: nil    类型为: *sync.Pool
	*/

	//var power int = 9000
	power := 9000 // :=，go可以推断变量的类型 (没有var 也没有类型关键字)
	// name, power := "Goku", 9000 //多个变量同时赋值

	fmt.Println("It's over %d\n", power)
}
