package main

import "fmt"

type myPerson struct {
	name string
	age  int8
}

// 给 myPerson 定义一个方法
// 值类型接受者
func (p myPerson) printInfo() {
	fmt.Printf("姓名:%v 年龄:%v \n", p.name, p.age)

}

// 值类型接受者 不会改变实例的类
func (p myPerson) setNameAndAge(name string, age int8) {
	p.name = name
	p.age = age
}

// 指针类型接收者 会修改实例的值
func (this *myPerson) setInfo(name string, age int8) {
	this.name = name
	this.age = age
}

func main() {
	p1 := myPerson{name: "小王子", age: 18}

	p1.printInfo() // 姓名:小王子 年龄:25
	p1.setNameAndAge("小公主", 18)
	p1.printInfo() // 姓名:小王子 年龄:18

	p1.setInfo("小公主", 18)
	p1.printInfo() // 姓名:小公主 年龄:18

}
