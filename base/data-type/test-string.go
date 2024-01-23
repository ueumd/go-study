package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
https://blog.csdn.net/m0_53732536/article/details/121776818
*/

func main() {
	str := "ABCDEFG"

	// Count strings.Count(s, sep string) int: 返回字符串s中有几个不重复的sep子串。
	fmt.Println(strings.Count(str, "")) // 8

	fmt.Println(utf8.RuneCountInString(str)) // 7
	fmt.Println(len(str))                    // 7

	str2 := "hello,世界"

	fmt.Println(strings.Count(str2, ""))      // 9
	fmt.Println(utf8.RuneCountInString(str2)) // 8
	fmt.Println(len(str2))                    // 12 汉字

	// 子串substr在s中，返回true
	fmt.Println(strings.Contains("team", "am")) //true

	// chars中任何一个Unicode代码点在s中，返回true
	fmt.Println(strings.ContainsAny("team", "a & m")) //true
}
