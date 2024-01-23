package main

import (
	"fmt"
	"strings"
)

func map1() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100

	// 如果 key 存在 ok 为 true, value 为对应的值
	// 不存在 ok 为 false,v value 为值类型的零值

	value, ok := scoreMap["张三"]

	if ok {
		fmt.Println(ok, value) // true 90
	} else {
		fmt.Println("查无此人")
	}

	value, ok = scoreMap["张大"]
	if ok {
		fmt.Println(ok, value)
	} else {
		fmt.Println("查无此人", ok, value) // 查无此人 false 0
	}

	// 删除
	scoreMap2 := make(map[string]int)
	scoreMap2["张三"] = 90
	scoreMap2["小明"] = 100
	scoreMap2["娜扎"] = 60
	delete(scoreMap2, "小明") //将小明:100 从 map 中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
		//张三 90
		//小明 100

	}
}

/*
*
元素为 map 类型的切片
*/
func map2() {
	// 元素为 map 类型的切片 [] map[type][type]
	var mapSlice = make([]map[string]string, 3)

	for index, value := range mapSlice {
		fmt.Printf("index: %d value: %v\n", index, value)
	}

	/**
	index: 0 value: map[]
	index: 1 value: map[]
	index: 2 value: map[]
	*/

	fmt.Println("after init \n")

	// 对切片中的 map 元素进行初始化
	mapSlice[0] = make(map[string]string, 10)

	mapSlice[0]["name"] = "Vue"
	mapSlice[0]["author"] = "Evan You"

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	/**
	index:0 value:map[author:Evan You name:Vue]
	index:1 value:map[]
	index:2 value:map[]

	*/
}

/*
*
值为切片类型的 map
*/
func map3() {
	// 注意值为切换 key: [...]
	// map[string][]string
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap) //map[]

	key := "中国"

	value, ok := sliceMap[key]
	if !ok {
		// 值为切片类型的 map
		value = make([]string, 0, 2)

	}

	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap) // map[中国:[北京 上海]]

	fmt.Println(sliceMap[key]) // [北京 上海]

}

func main() {
	var wordMap = make(map[string]int)
	var str = "how do you do"

	// 字符串转数组
	var arrSlice = strings.Split(str, " ") // 按空格转数组
	fmt.Println(arrSlice)

	for _, word := range arrSlice {
		wordMap[word]++
	}

	fmt.Println(wordMap) // map[do:2 how:1 you:1]

}
