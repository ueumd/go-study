package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func change(u User) {
	u.Age = 21
	fmt.Println(u.Age)
}

func change2(u *User) {
	u.Age = 18
	fmt.Println(u.Age)
}

func (u User) changeAge() {
	u.Age = 22
	fmt.Println(u.Age)
}

func (u *User) changeAgePointer() {
	u.Age = 100
	fmt.Println(u.Age)
}

func ptest() {
	var age *int //声明指针变量
	num := 100   //声明实际变量
	age = &num   //指针变量的存储地址

	fmt.Println(&num) //0xc042052080
	fmt.Println(age)  //0xc042052080
	fmt.Println(*age) //100

	*age = 1000
	fmt.Println(*age) //1000
}

func testuser() {
	u := &User{
		Name: "小明",
		Age:  18,
	}

	fmt.Println(u) //&{小明 18}
	u.Age = 19
	fmt.Println(u) //&{小明 19}

	change(*u) //21
	change2(u) //18

	/**
	change3(&u)
	change4(&u)
	这样每次都要取地址符号，所以一般在初始化时把地取出来 &User{}
	u := &User{
		Name:"小明",
		Age:18,
	}
	change3(u)
	change4(u)
	*/

	fmt.Println(u) //&{小明 18}

	u.changeAge()        //22
	fmt.Println(u)       //&{小明 19}
	u.changeAgePointer() //100
	fmt.Println(u)       //&{小明 100}

}

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func anonymous() {
	//匿名结构体
	p := &struct {
		Name string
		Age  int
	}{
		Name: "小王",
		Age:  18,
	}

	fmt.Println(p) //&{小王 18}
}

func (p *person) changeAgeP() {
	p.Age = 100
	fmt.Println(p.Age)
}

func testp() {
	p := person{Name: "joe", Age: 19}
	p.Contact.Phone = "13900000000"
	p.Contact.City = "上海"
	fmt.Println(p) //{joe 19 {13900000000 上海}}
	p.changeAgeP() //100
	fmt.Println(p) //{joe 100 {13900000000 上海}}
}

/*
*
声明一个底层类型为int的类型，并实现调用某个方法就递增100
a:=0 调用a.Increase() 变成100
*/
type TZ int

func (tz *TZ) Increase(num int) {
	*tz += TZ(num) //强制转换成TZ类型
}

func testTZ() {
	var a TZ
	a.Increase(100)
	fmt.Println(a) //100
}

func init() {
	//ptest()
	testuser()
	//anonymous()
	//testp()
	testTZ()
}
