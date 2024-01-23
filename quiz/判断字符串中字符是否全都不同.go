package main

import (
	"fmt"
	"strings"
)

/*
*
判断字符串中字符是否全都不同
https://zhuanlan.zhihu.com/p/375388960
https://www.getcoder.cn/archives/-ji-chu-2-pan-duan-zi-fu-chuan-zhong-zi-fu-shi-fou-quan-dou-bu-tong

问题描述
请实现⼀个算法，确定⼀个字符串的所有字符是否全都不同。这⾥我们要求不允许使⽤额外的存储结构。

给定⼀个string，请返回⼀个bool值，true代表所有字符全都不同，false代表存在相同的字符。

保证字符串中的字符为ASCII字符。

字符串的⻓度⼩于等于3000
*/
func isUniqueString(str string) bool {
	if strings.Count(str, "") > 3000 {
		return false
	}

	for key, value := range str {
		// ASCII字符字符一共有256个，其中128个是常用字符，可以在键盘上输入。128之后的是键盘上无法找到的
		if value > 127 {
			return false
		}

		// Count strings.Count(s, sep string) int: 返回字符串s中有几个不重复的sep子串。
		if strings.Count(str, string(value)) > 1 {
			return false
		}

		// 或者
		if strings.Index(str, string(value)) != key {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isUniqueString("Helxo@afdgy"))
}
