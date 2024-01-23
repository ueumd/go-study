package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/**
https://www.yuque.com/ueumd/blog/ghabuf
结体构 与 tag
 */
type PersonInfo struct {
	Name string	`orm:name,max_length=17,min_length=5`
	Age int 	`json:"age,omitempty"` // omitempty 在序列化json会忽略该字段的默认值
	Gender string `json:"gender"`
	Salary string `json:"-"` // 不会被序列化
}


type PersonInfo2 struct {
	Name string	`orm:"name, max_length=17, min_length=5""`
}

func main()  {
	info := &PersonInfo{
		Name:"Tom",
		Gender: "猫",
		Salary: "1k",
	}

	re, _ := json.Marshal(info)

	// Age 没有声明，会默认值，但在tag里采用了 omitempty 在序列化时会忽略该字段
	fmt.Println(string(re)) // {"name":"Tom","Gender":"猫"}


	info2 := PersonInfo2{
		"Jerry",
	}
	// 自己加标签取出来 通过反射包
	t := reflect.TypeOf(info2)

	fmt.Println("Type", t.Name())
	fmt.Println("Kind", t.Kind())

	for i:= 0; i<t.NumField(); i++ {
		filed := t.Field(i)
		tag := filed.Tag.Get("orm")
		fmt.Println(tag) // name, max_length=17, min_length=5

		// 再利用字符串处理可以提取key value
		fmt.Printf("%d, %v (%v), tag:'%v'\n", i+1, filed.Name, filed.Type.Name(), tag) // 1, Name (string), tag:'name, max_length=17, min_length=5'

	}
}