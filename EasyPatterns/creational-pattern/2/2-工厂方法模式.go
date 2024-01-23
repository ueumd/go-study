package main

import (
	"fmt"
)

// ===================== 抽象层 ==============================

// 水果类(抽象接口)
type Fruit2 interface {
	show() //接口的某方法
}

// 工厂类(抽象接口)
type AbstractFactory2 interface {
	createFruit() Fruit2 //生产水果类(抽象)的生产器方法
}

// ===================== 基础类模块 ==============================
type Apple2 struct {
}

func (apple *Apple2) show() {
	//TODO implement me
	panic("implement me")
}

type Banana2 struct {
}

func (banana *Banana2) show() {
	fmt.Println("我是香蕉")
}

type Pear2 struct {
}

func (pear *Pear2) show() {
	fmt.Println("我是梨")
}

// ===================== 工厂模块 ==============================
// 具体的苹果工厂
type AppleFactory2 struct {
}

func (this *AppleFactory2) createFruit() Fruit2 {
	var fruit Fruit2

	//生产一个具体的苹果
	fruit = new(Apple2)

	return fruit
}

// 具体的香蕉工厂
type BananaFactory2 struct {
}

func (this *BananaFactory2) createFruit() Fruit2 {
	var fruit Fruit2

	//生产一个具体的苹果
	fruit = new(Banana2)

	return fruit
}

// 具体的梨工厂
type PearFactory2 struct {
}

func (this *PearFactory2) createFruit() Fruit2 {
	var fruit Fruit2

	//生产一个具体的苹果
	fruit = new(Pear2)

	return fruit
}

// ===================== 工厂模块 ==============================
func main() {
	/*
		本案例为了突出根据依赖倒转原则与面向接口编程特性。
		一些变量的定义将使用显示类型声明方式
	*/

	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂

	var appleFac AbstractFactory2
	appleFac = new(AppleFactory2)

	var apple Fruit2
	apple = appleFac.createFruit()
	apple.show()

	//需求2：需要一个具体的香蕉对象
	//1-先要一个具体的香蕉工厂
	var bananaFac AbstractFactory2
	bananaFac = new(BananaFactory2)

	//2-生产相对应的具体水果
	var banana Fruit2
	banana = bananaFac.createFruit()

	banana.show()

	//需求3：需要一个具体的梨对象
	//1-先要一个具体的梨工厂
	var pearFac AbstractFactory2
	pearFac = new(PearFactory2)
	//2-生产相对应的具体水果
	var pear Fruit2
	pear = pearFac.createFruit()

	pear.show()

	//需求4：需要一个日本的苹果？
	//1-先要一个具体的日本评估工厂
	var japanAppleFac AbstractFactory2
	japanAppleFac = new(JapanAppleFactory2)
	//2-生产相对应的具体水果
	var japanApple Fruit2
	japanApple = japanAppleFac.createFruit()

	japanApple.show()
}

/*
新增的基本类“日本苹果”，和“具体的工厂” 均没有改动之前的任何代码。完全符合开闭原则思想。
新增的功能不会影响到之前的已有的系统稳定性。

优点：
1. 不需要记住具体类名，甚至连具体参数都不用记忆。
2. 实现了对象创建和使用的分离。
3. 系统的可扩展性也就变得非常好，无需修改接口和原类。
4.对于新产品的创建，符合开闭原则。

缺点：
1. 增加系统中类的个数，复杂度和理解度增加。
2. 增加了系统的抽象性和理解难度。

适用场景：
1. 客户端不知道它所需要的对象的类。
2. 抽象工厂类通过其子类来指定创建哪个对象

*/

// (+) 新增一个"日本苹果"
type JapanApple2 struct {
}

func (jp *JapanApple2) show() {
	fmt.Println("我是日本苹果")
}

// 具体的日本工厂
type JapanAppleFactory2 struct {
	AbstractFactory2
}

func (fac *JapanAppleFactory2) createFruit() Fruit2 {
	var fruit Fruit2

	//生产一个具体的日本苹果
	fruit = new(JapanApple2)

	return fruit
}
