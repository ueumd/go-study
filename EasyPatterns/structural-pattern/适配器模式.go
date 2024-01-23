package main

import "fmt"

// 适配的目标
type V5 interface {
	UseV5()
}

// 业务类，依赖V5接口
type Phone2 struct {
	v V5
}

func NewPhone(v V5) *Phone2 {
	return &Phone2{}
}

func (p *Phone2) Charge() {
	fmt.Println("Phone进行充电...")
	p.v.UseV5()
}

// 被适配的角色，适配者
type V220 struct {
}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) UseV5() {
	fmt.Println("使用适配器进行充电")
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

// ------- 业务逻辑层 -------
func main() {
	iphone := NewPhone(NewAdapter(new(V220)))

	iphone.Charge()
}
