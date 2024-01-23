package main

import "fmt"

/*
合成复用原则
如果使用继承，会导致父类的任何变换都可能影响到子类的行为。

如果使用对象组合，就降低了这种依赖关系。对于继承和组合，优先使用组合

*/

type cat struct {
}

func (this *cat) eat() {
	fmt.Println("小猫吃饭")
}

/*===================== 使用继承来实现 ===========================================*/
// 给小猫添加一个 可以睡觉的方法 （使用继承来实现）
type catB struct {
	cat // 继承
}

func (this *catB) sleep() {
	fmt.Println("小猫睡觉")
}

/*===================== 使用组合的方式 ===========================================*/
// 使用组合的方式
// 给小猫添加一个 可以睡觉的方法 （使用组合的方式）
type catC struct {
	C *cat // 组合
}

func (this *catC) sleep() {
	fmt.Println("小猫睡觉")
}

func (this *catC) eat() {
	this.C.eat()
}

// 继承会把父类所有方法拿过来，如果父类再有父类，耦合度太高了
// 组合不会把有父类所有的方法都拿过来了
// 组合只是把某个方法拿过来使用
func (this *catC) eat2(c *cat) {
	// this.C.eat()
	c.eat()
}

func main() {

	c := &cat{}
	c.eat()

	println("--------------------")

	//通过继承增加的功能，可以正常使用
	cb := new(catB)
	cb.eat() // 继承了父类的方法
	cb.sleep()

	println("--------------------")

	//通过组合增加的功能，可以正常使用
	cc := new(catC)

	cc.C = new(cat)
	cc.C.eat()

	// cc.C.sleep() // 没有sleep方法

	cc.eat()
	cc.sleep()

	cc.eat2(c)

}
