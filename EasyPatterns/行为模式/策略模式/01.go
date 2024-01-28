package main

import "fmt"

/**
策略（Strategy）模式的定义：该模式定义了一系列算法，并将每个算法封装起来，使它们可以相互替换，且算法的变化不会影响使用算法的客户。
策略模式属于对象行为模式，它通过对算法进行封装，把使用算法的责任和算法的实现分割开来，并委派给不同的对象对这些算法进行管理。

策略模式的主要优点如下。

多重条件语句不易维护，而使用策略模式可以避免使用多重条件语句，如 if...else 语句、switch...case 语句。
策略模式提供了一系列的可供重用的算法族，恰当使用继承可以把算法族的公共代码转移到父类里面，从而避免重复的代码。
策略模式可以提供相同行为的不同实现，客户可以根据不同时间或空间要求选择不同的。
策略模式提供了对开闭原则的完美支持，可以在不修改原代码的情况下，灵活增加新算法。
策略模式把算法的使用放到环境类中，而算法的实现移到具体策略类中，实现了二者的分离。
其主要缺点如下。

客户端必须理解所有策略算法的区别，以便适时选择恰当的算法类。
策略模式造成很多的策略类，增加维护难度。

https://developer.aliyun.com/article/1268033
https://zhuanlan.zhihu.com/p/391710236

策略模式属于行为型设计模式，它允许在运行时选择算法的行为。策略模式主要包含以下角色：
环境（Context）：定义客户端所感兴趣的接口，并维护一个具体策略的实例。
抽象策略（Strategy）：定义一个接口，用于封装具体策略的行为。
具体策略（Concrete Strategy）：实现抽象策略定义的接口，完成具体策略对应的行为。
*/

/**
示例场景：
为了更好地理解策略模式的应用，我们以一个简单的示例场景为例：
假设我们正在开发一个电商系统，针对不同的用户类型（普通用户、VIP用户、超级VIP用户），我们希望根据用户类型来计算商品的折扣价格。
*/

// -------------------------------------环境接口

// 环境（PricingContext）接口定义了客户端所感兴趣的接口，以及维护具体策略的方法。
type PricingContext interface {
	// 设置策略方法
	SetStrategy(strategy PricingStrategy)
	CalculatePrice(originalPrice float64) float64
}

// ------------------------------------- 抽象策略接口
// 接口定义了策略的行为
type PricingStrategy interface {
	CalculatePrice(originalPrice float64) float64
}

// ------------------------------------- 具体策略：普通用户策略

// 具体策略（RegularUserStrategy、VipUserStrategy、SuperVipUserStrategy）实现了抽象策略接口，分别定义了普通用户、VIP用户和超级VIP用户的策略。
type RegularUserStrategy struct{}

// 具体策略：普通用户策略
func (s *RegularUserStrategy) CalculatePrice(originalPrice float64) float64 {
	return originalPrice
}

// 具体策略：VIP用户策略
type VipUserStrategy struct{}

func (s *VipUserStrategy) CalculatePrice(originalPrice float64) float64 {

	return originalPrice * 0.9
}

// 具体策略：超级VIP用户策略
type SuperVipUserStrategy struct{}

func (s *SuperVipUserStrategy) CalculatePrice(originalPrice float64) float64 {
	return originalPrice * 0.8
}

// ------------------------------------- 具体环境

// 具体环境（PricingService）维护当前策略，并在需要计算价格时调用具体策略的方法。
type PricingService struct {
	strategy PricingStrategy
}

func (s *PricingService) SetStrategy(strategy PricingStrategy) {
	s.strategy = strategy
}

func (s *PricingService) CalculatePrice(originalPrice float64) float64 {
	return s.strategy.CalculatePrice(originalPrice)
}

func main() {
	pricingService := &PricingService{}
	originalPrice := 100.0

	// 普通用户
	pricingService.SetStrategy(&RegularUserStrategy{})
	discountedPrice := pricingService.CalculatePrice(originalPrice)
	fmt.Printf("普通用户原价：%v，折扣价：%v\n", originalPrice, discountedPrice)

	// vip用户
	pricingService.SetStrategy(&VipUserStrategy{})
	discountedPrice = pricingService.CalculatePrice(originalPrice)
	fmt.Printf("VIP用户原价：%v，折扣价：%v\n", originalPrice, discountedPrice)

	// 超级VIP用户
	pricingService.SetStrategy(&SuperVipUserStrategy{})
	discountedPrice = pricingService.CalculatePrice(originalPrice)
	fmt.Printf("超级VIP用户原价：%v，折扣价：%v\n", originalPrice, discountedPrice)
}
