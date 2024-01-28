package main

import "fmt"

/*

 外观模式中角色和职责
Façade(外观角色)：为调用方, 定义简单的调用接口。
SubSystem(子系统角色)：功能提供者。指提供功能的类群（模块或子系统）。

优点：
(1) 它对客户端屏蔽了子系统组件，减少了客户端所需处理的对象数目，并使得子系统使用起来更加容易。通过引入外观模式，客户端代码将变得很简单，与之关联的对象也很少。
(2) 它实现了子系统与客户端之间的松耦合关系，这使得子系统的变化不会影响到调用它的客户端，只需要调整外观类即可。
(3) 一个子系统的修改对其他子系统没有任何影响。
缺点：
(1) 不能很好地限制客户端直接使用子系统类，如果对客户端访问子系统类做太多的限制则减少了可变性和灵活 性。
(2) 如果设计不当，增加新的子系统可能需要修改外观类的源代码，违背了开闭原则。

4.4.5 适用场景
(1) 复杂系统需要简单入口使用。
(2) 客户端程序与多个子系统之间存在很大的依赖性。
(3) 在层次化结构中，可以使用外观模式定义系统中每一层的入口，层与层之间不直接产生联系，而通过外观类建立联系，降低层之间的耦合度。
*/

type SubSystemA struct {
}

func (sa *SubSystemA) MethodA() {
	fmt.Println("子系统方法A")
}

type SubSystemB struct {
}

func (sb *SubSystemB) MethodB() {
	fmt.Println("子系统方法B")
}

type SubSystemC struct {
}

func (sc *SubSystemC) MethodC() {
	fmt.Println("MethodC")
}

type SubSystemD struct{}

func (sd *SubSystemD) MethodD() {
	fmt.Println("子系统方法D")
}

//外观模式，提供了一个外观类， 简化成一个简单的接口供使用

type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *Facade) MethodTwo() {
	f.c.MethodC()
	f.d.MethodD()
}

func main() {
	// 不使用外观
	sa := new(SubSystemA)
	sa.MethodA()

	sb := new(SubSystemB)
	sb.MethodB()

	fmt.Println("-----------")
	//使用外观模式
	f := Facade{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
		d: new(SubSystemD),
	}

	// 调用外观模式
	f.MethodOne()
}
