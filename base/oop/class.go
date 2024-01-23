package main

import "fmt"

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

/*
*
不支持构造函数、析构函数
可以通过定义形如 NewXXX 这样的全局函数（首字母大写）作为类的初始化函数：
*/
func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

// 指定字段初始化
func NewStudent2(id uint, name string, score float64) *Student {
	return &Student{id: id, name: name, score: score}
}

/*
*
由于不需要对类的成员变量进行修改，所以不需要传入指针，
而 SetXXX 方法需要在函数内部修改成员变量的值，并且该修改要作用到该函数作用域以外所以需要传入指针类型（结构体是值类型，不是引用类型，所以需要显式传入指针）。
*/
func (s Student) GetName() string {
	return s.name
}

/*
*
值方法 : 不能修改实例成员属性
*/
func (s Student) SetName(name string) {
	s.name = name
}

/*
*
指针方法
*/
func (s *Student) SetNewName(name string) {
	s.name = name
}

func main() {
	tom := NewStudent2(1, "Tom", 100)
	jack := NewStudent2(2, "Jack", 1000)
	fmt.Println(tom)  // &{1 Tom false 100}
	fmt.Println(jack) // &{2 Jack false 1000}

	// 值方法
	tom.SetName("Jack")
	jack.SetName("Tom")

	fmt.Println(tom)  // &{1 Tom false 100}
	fmt.Println(jack) // &{2 Jack false 100}

	fmt.Println(tom.GetName())
	fmt.Println(jack.GetName())

	// 指针方法
	tom.SetNewName("Jack")
	jack.SetNewName("Tom")

	fmt.Println(tom)  //&{1 Jack false 100}
	fmt.Println(jack) //&{2 Tom false 1000}

	fmt.Println(tom.GetName())
	fmt.Println(jack.GetName())

}

/**
值方法和指针方法的区别

在 Go 语言中，当我们将成员方法 SetName 所属的类型声明为指针类型时，严格来说，该方法并不属于 Student 类，而是属于指向 Student 的指针类型
归属于 Student 的成员方法只是  Student 类型下所有可用成员方法的子集
归属于 *Student 的成员方法才是 Student 类完整可用方法的集合

调用方法时，之所以可以直接在 student 实例上调用 SetName 方法，是因为 Go 语言底层会自动将 student 转化为对应的指针类型 &student
所以真正调用的代码是 (&student).SetName("Tom")
*/
