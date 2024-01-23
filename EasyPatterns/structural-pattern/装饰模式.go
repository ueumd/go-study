package main

import "fmt"

// --------------------------------- 抽象层 ------------------------------------

// 抽象的构件
type Phone interface {
	show() // 构建的功能
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) show() {}

// ----------------------------- 实现层 ------------------------------------------
// 具体的构建

// ---------------------   手机    --------------------------
type Huawei struct{}

func (hw *Huawei) show() {
	fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct{}

func (xm *XiaoMi) Show() {
	fmt.Println("秀出了XiaoMi手机")
}

// -----------------------   装饰    --------------------------
// 具体的装饰器类 手机膜
type MoDecorator struct {
	Decorator //继承基础装饰器类(主要继承Phone成员属性)
}

func (md *MoDecorator) show() {
	md.phone.show()      //调用被装饰构件的原方法
	fmt.Println("贴膜的手机") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

// 具体的装饰器类 手机壳
type KeDecorator struct {
	Decorator //继承基础装饰器类(主要继承Phone成员属性)
}

func (md *KeDecorator) show() {
	md.phone.show()       //调用被装饰构件的原方法
	fmt.Println("手机壳的手机") //装饰额外的方法
}

func newKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

// --------------------------------------- 业务逻辑层 -------------------------------

func main() {
	var huawei Phone
	huawei = new(Huawei)
	huawei.show()

	println("\n--------------------")

	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei)
	moHuawei.show()

	println("\n--------------------")

	var keHuawei Phone
	keHuawei = newKeDecorator(huawei)
	keHuawei.show()

	println("\n--------------------")

	var keMoHuaWei Phone
	keMoHuaWei = NewMoDecorator(keHuawei)
	keMoHuaWei.show()
}
