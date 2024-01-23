package main

import "fmt"

/**


map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取。
map的长度是不固定的，也就是和slice一样，也是一种引用类型。
内置的len函数同样适用于map，返回map拥有的key的数量。
map的值可以很方便的修改，通过重新赋值即可。
*/

// map初始化
func init() {
	type personInfo struct {
		ID      string
		Name    string
		Address string
	}

	//1 直接初始化（创建）
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	myMap := map[string]personInfo{"1234": personInfo{"1", "Jack", "Room 101,..."}}

	value, key := rating["C"] //判断key "C"是否存在，存在key C就打印
	if key {
		fmt.Println(value) //5
	}

	v, k := myMap["1234"]
	if k {
		fmt.Println(v)      //{1 Jack Room 101,...}
		fmt.Println(v.Name) //Jack
	}

	// 2通过make初始化（创建）
	numbers := make(map[string]int, 3)
	//map元素赋值
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3

	//上面方式的简写方法
	numbers2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for nk, nv := range numbers2 {
		fmt.Println(nk, ":", nv)
		if nv == 2 {
			fmt.Println("haha")
		}
		/**
		one : 1
		two : 2
		haha
		three : 3
		*/
	}

}
