package main

import "fmt"

// SAddress 地址结构体
type SAddress struct {
	Province string
	City     string
}

// SUser 用户结构体
type SUser struct {
	Name   string
	Gender string

	//匿名结构体 SAddress 只有类型，没有名字，匿名字段，继承了SAddress的成员
	SAddress
}

func main() {
	var user SUser

	user.Name = "小王子"
	user.Gender = "男"
	user.SAddress.Province = "上海" //通过匿名结构体.字段名访问
	user.City = "上海市"             //直接访问匿名结构体的字段名

	fmt.Printf("user = %#v\n", user)
	//user = main.SUser{Name:"小王子", Gender:"男", SAddress:main.SAddress{Province:" 上海", City:"上海市"}}

	fmt.Printf("user = %+v\n", user) // user = {Name:小王子 Gender:男 SAddress:{Province:上海 City:上海市}}

}
