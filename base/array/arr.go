package main

import (
	"fmt"
)

func arr1() {
	var a = []string{"A", "B", "C"}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i]) //ABC
	}
	fmt.Println("")
	b := [6]int{3, 6, 9}
	for i := 0; i < len(b); i++ {
		fmt.Print(b[i]) //369000 自动补0
	}

	fmt.Println()
	var a1 [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	/**
	range遍历数组，切片和字典
	数组和切片 返回索引和元素；
	字典		  返回键和值
	*/
	for row, v := range a1 {
		for col, v1 := range v {
			fmt.Printf("(%d, %d)=%d ", row, col, v1)
		}
		fmt.Println()
	}

}

func map11111() {
	/**
	var a map[keytype]valuetype
	map声明是不会分配内存的，需要make初始化。
	*/

	m := make(map[string]string, 10)
	m["name"] = "Nick"
	m["age"] = "20"

	fmt.Println(m["name"])

}

func init1() {
	//arr1()
	map11111()
}
