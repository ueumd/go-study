package main

import (
	"fmt"
	"strconv"
)

type User struct {
	name string
	age  int
	sex  string

	phone string
}

func (s *User) SetName(name string) {
	s.name = name
}

func (s *User) GetName() string {
	return s.name
}

func (s *User) SetAge(age int) {
	s.age = age
}

func (s *User) GetAge() int {
	return s.age
}

func (s *User) String() string {
	return "name is " + s.name + ",age is " + strconv.Itoa(s.age) + " ,sex=" + s.sex + " ,phone=" + s.phone
}

func (s *User) SetSex(sex string) {
	s.sex = sex
}

func (s *User) GetSex() string {
	return s.sex
}

func (s *User) SetPhone(phone string) {
	s.phone = phone
}

func (s *User) GetPhone() string {
	return s.phone
}

func (User) Print() string {
	return "print"
}

type person struct {
	ID      string
	Name    string
	Address string
}

func map1() {

	// map声明是不会分配内存的，需要make初始化。var a map[keytype]valuetype
	userDB := make(map[string](*User)) //注意写法 [keytype]valuetype

	//初始化，注意对数组的初始化

	u := new(User)
	u.SetAge(12)
	u.SetName("张三")
	u.SetSex("男")
	u.SetPhone("15902783102")

	userDB["u1"] = u

	v, ok := userDB["u1"] //获取map值  key为'u1'
	fmt.Println(ok)
	if !ok {
		fmt.Println(" 没有找到信息")
		return
	}

	//打印出全部值 和各个分值
	fmt.Println(v.String())
	fmt.Printf("userDB[u1] = \n {name=%v \n age=%v \n sex=%v \n phone=%v \n}", v.GetName(), v.GetAge(), v.GetSex(), v.GetPhone())

}

func init() {
	map1()
}

//https://studygolang.com/articles/2379
//https://studygolang.com/articles/2494
