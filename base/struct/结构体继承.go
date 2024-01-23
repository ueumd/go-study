package main

import "fmt"

type PersonStruct struct {
	name string // 名字
	sex  string // 性别, 字符类型
	age  int    // 年龄
}

func (p *PersonStruct) showName() {
	fmt.Printf("我的名子叫: %s \n", p.name)
}

/*
*
结构体的继承 - 结构体指针类型
通过嵌套匿名结构体实现继承
*/
type StudentStruct struct {
	*PersonStruct //通过嵌套匿名结构体实现继承
	id            int
	addr          string
}

func (s *StudentStruct) study() {
	fmt.Printf("%s 同学正在学习 \n", s.name)
}

func main() {

	// 初始化 指针类型取址 注意嵌套的是结构体指针
	s1 := StudentStruct{&PersonStruct{"小王子", "男", 18}, 1, "上海"}

	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr) // 小王子 男 18 1 上海
	s1.showName()                                        // 我的名子叫: 小王子
	s1.study()                                           // 小王子 同学正在学习

	s2 := &StudentStruct{
		id:   2,
		addr: "北京",
		PersonStruct: &PersonStruct{
			name: "小公主",
			age:  18,
			sex:  "女",
		},
	}

	s2.showName() // 我的名子叫: 小公主
	s2.study()    // 小公主 同学正在学习

	var ss = new(PersonStruct)
	fmt.Println(ss.name == "")
	fmt.Println(ss)

}
