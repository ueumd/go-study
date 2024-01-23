package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Call() string {
	return "动物的叫声..."
}

func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}

func (a Animal) GetName() string {
	return a.Name
}

type Dog struct {
	Animal
}

// 子类 重写 父类方法
func (d Dog) Call() string {
	return "汪汪汪"
}

func (a Dog) FavorFood() string {
	return "骨头..."
}

func main() {
	animal := Animal{"中华田园犬"}

	dog := Dog{
		animal,
	}

	fmt.Println(dog.GetName())

	// 子类 重写 父类方法，子类调用自己的方法，不重写则调用父类的方法
	fmt.Println(dog.Call())
	fmt.Println(dog.FavorFood())

	// go 没有专门提供引用父类实例的关键字 super、parent
	// 如果要继续调用父类方法可以像下面这样
	fmt.Println(dog.Animal.Call())
	fmt.Println(dog.Animal.FavorFood())
}
