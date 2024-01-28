package main

import "fmt"

/*
5.1.3 模板方法的优缺点
优点：
(1) 在父类中形式化地定义一个算法，而由它的子类来实现细节的处理，在子类实现详细的处理算法时并不会改变算法中步骤的执行次序。
(2) 模板方法模式是一种代码复用技术，它在类库设计中尤为重要，它提取了类库中的公共行为，将公共行为放在父类中，而通过其子类来实现不同的行为，它鼓励我们恰当使用继承来实现代码复用。
(3) 可实现一种反向控制结构，通过子类覆盖父类的钩子方法来决定某一特定步骤是否需要执行。
(4) 在模板方法模式中可以通过子类来覆盖父类的基本方法，不同的子类可以提供基本方法的不同实现，更换和增加新的子类很方便，符合单一职责原则和开闭原则。

缺点：
需要为每一个基本方法的不同实现提供一个子类，如果父类中可变的基本方法太多，将会导致类的个数增加，系统更加庞大，设计也更加抽象。
*/

// 抽象类，制作饮料,包裹一个模板的全部实现步骤
type Beverage interface {
	BoilWater() //煮开水
	Brew()      // 冲泡
	PourInCup() //倒入杯中
	AddThings() //添加酌料

	WantAddThings() bool //是否加入酌料Hook
}

type template struct {
	b Beverage
}

// 封装的固定模板
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	//子类可以重写该方法来决定是否执行下面动作
	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

// 具体的模板子类 制作咖啡
type MakeCoffee struct {
	template //继承模板
}

func (m MakeCoffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

func (m MakeCoffee) Brew() {
	fmt.Println("用水冲咖啡豆")
}

func (m MakeCoffee) PourInCup() {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}

func (m MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (m MakeCoffee) WantAddThings() bool {
	return true //启动Hook条件
}

func NewMakeCaffe() *MakeCoffee {
	makeCaffe := new(MakeCoffee)

	//b 为Beverage，是MakeCaffee的接口，这里需要给接口赋值，指向具体的子类对象
	//来触发b全部接口方法的多态特性。
	makeCaffe.b = makeCaffe
	return makeCaffe
}

// 具体的模板子类 制作茶
type MakeTea struct {
	template // 继承模版
}

func (m MakeTea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

func (m MakeTea) Brew() {
	fmt.Println("用水冲茶叶")
}

func (m MakeTea) PourInCup() {
	fmt.Println("将充好的咖啡倒入茶壶中")
}

func (m MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (m MakeTea) WantAddThings() bool {
	return false //关闭Hook条件
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func main() {
	makeCoffee := NewMakeCaffe()
	makeCoffee.MakeBeverage()

	fmt.Println("------------------- \n")

	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}
