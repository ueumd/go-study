package main

import "fmt"

/**
单一职责原则
类的职责单一，对外只提供一种功能，而引起类的变化的原因都应该只有一个
*/

/*
不单一
*/
// 穿衣服的方式
type clothes struct{}

func (c *clothes) onWork() {
	fmt.Println("工作的装扮")
}

func (c *clothes) onShop() {
	fmt.Println("工作的装扮")
	// fmt.Println("逛街的装扮")
}

func test() {
	// 不单一个，一个类有多种功能
	c := &clothes{}
	// 工作的业务
	c.onWork()

	// 逛街的业务
	c.onShop()
}

/*===============================================================*/
type clothesShop struct{}

func (c *clothesShop) style() {
	fmt.Println("逛街的装扮")
}

type clothesWork struct{}

func (c *clothesWork) style() {
	fmt.Println("工作的装扮")
}

func main() {
	// 类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。

	// 在面向对向过程中，设计一个类，建议对外提供的功能单一，接口单一，影响一个类的范围就只限定在这一个接口上。
	// 一个类的一个接口具备这个类的功能含义，职责单一不复杂

	// 工作业务
	cw := &clothesWork{}
	cw.style()

	// 逛街业务
	cs := &clothesShop{}
	cs.style()
}
