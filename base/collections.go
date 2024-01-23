package main

import (
	"fmt"
)

func MyArr() {
	// arr :=[5]byte {'A', 'B', 'C', 'D', 'E'}
	// 动态数组 自动分配内存空间
	slice := []byte{'A', 'B', 'C', 'D', 'E'}

	for key, val := range slice {
		//key val
		//0 65
		//1 66
		//2 67
		//3 68
		//4 69
		// Println没有格式化的功能
		fmt.Printf("[%d]=%c \n", key, val) // [0]=A [1]=B [2]=C [3]=D [4]=E
	}
}
