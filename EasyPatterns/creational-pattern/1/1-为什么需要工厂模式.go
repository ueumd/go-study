package main

import (
	"fmt"
)

// 水果类
type Fruit struct {
	//......
}

func (this *Fruit) showName(name string) {
	if name == "apple" {
		fmt.Println("我是苹果")
	} else if name == "banana" {
		fmt.Println("我是香蕉")
	} else if name == "pear" {
		fmt.Println("我是梨")
	}
}

/*
Fruit类的职责过重
1. 它负责初始化和显示所有的水果对象，将各种水果对象的初始化代码和显示代码集中在一个类中实现，违反了“单一职责原则”，不利于类的重用和维护；
2. 当需要增加新类型的水果时，必须修改Fruit类的构造函数NewFruit()和其他相关方法源代码，违反了“开闭原则”。扩展而不是修改
*/
func newFruit(name string) *Fruit {
	fruit := new(Fruit)

	if name == "apple" {
		//创建apple逻辑
	} else if name == "banana" {
		//创建banana逻辑
	} else if name == "pear" {
		//创建pear逻辑
	}

	return fruit
}

func main() {
	apple := newFruit("apple")
	apple.showName("apple")

	banana := newFruit("banana")
	banana.showName("banana")

	pear := newFruit("pear")
	pear.showName("pear")
}
