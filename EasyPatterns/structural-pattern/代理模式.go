package main

import "fmt"

type Goods struct {
	Kind string // 商品种类
	Fact bool   // 商品真伪
}

// ======================================= 抽象层 ===============================
type Shopping interface {
	Buy(goods *Goods)
}

// ======================================= 实现层 ===============================

// 去韩国购物
type KoreaShopping struct{}

func (this *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物, 买了 ", goods.Kind)
}

// 去美国购物
type AmericanShopping struct{}

func (this *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物, 买了 ", goods.Kind)
}

// 去非洲购物
type AfricaShopping struct{}

func (as *AfricaShopping) Buy(goods *Goods) {
	fmt.Println("去非洲进行了购物, 买了 ", goods.Kind)
}

// 海外代理
type OverseasProxy struct {
	shopping Shopping
}

// 验货流程
func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

// 安检流程
func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1. 先验货
	if op.distinguish(goods) == true {
		//2. 进行购买

		//调用原被代理的具体主题任务
		op.shopping.Buy(goods)

		//3 海关安检
		op.check(goods)
	}
}

// 创建一个代理,并且配置关联被代理的主题
func newProxy(shopping Shopping) Shopping {
	return &OverseasProxy{
		shopping: shopping,
	}
}

func main() {
	goods1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	goods2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	var shopping Shopping

	//如果不使用代理来完成从韩国购买任务
	shopping = new(KoreaShopping)
	//1-先验货
	if goods1.Fact == true {
		fmt.Println("对[", goods1.Kind, "]进行了辨别真伪.")
		//2-去韩国购买
		shopping.Buy(&goods1)
		//3-海关安检
		fmt.Println("对[", goods1.Kind, "] 进行了海关检查， 成功的带回祖国")
	}

	fmt.Println("---------------以下是 使用 代理模式-------")

	var overseasProxy Shopping
	overseasProxy = newProxy(shopping)

	overseasProxy.Buy(&goods1)
	overseasProxy.Buy(&goods2)
}
