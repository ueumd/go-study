package main

import "fmt"

/*
开闭原则
https://www.yuque.com/aceld/lfhu8y/ckdht9
*/

// 平铺式设计

/*
一个银行业务员，他可能拥有很多的业务，比如Save()存款、Transfer()转账、Pay()支付等
*/
// Banker银行业务员
type banker struct {
}

// 存款业务
func (this *banker) save() {
	fmt.Println("进行了 存款业务...")
}

func (this *banker) transfer() {
	fmt.Println("进行了 转账业务...")
}

func (this *banker) pay() {
	fmt.Println("进行了 支付业务...")
}

func _main() {
	bankerOne := banker{}

	// 平铺式设计
	bankerOne.save()
	bankerOne.transfer()
	bankerOne.pay()

	// bankerOne.xx()
	// bankerOne.xx()
	// bankerOne.xx()

	// 假如此时Banker已经有99个业务了,现在我们要添加第100个业务
	// 可能由于一次不小心，导致之前99个业务全部崩溃
	// 这样所有的业务都在banker类里，耦合度太高，职责也不够单一
	// bankerOne.xx()
}

/*
开闭原则设计：
如果我们拥有接口，interface,那我们抽像一个Banker，然后再提供一个抽象方法
分别根据这个抽象模块，去实现支付Banker（实现支付方法）,转账Banker（实现转账方法）

开闭原则:一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。
简单的说就是在修改需求的时候，应该尽量通过扩展来实现变化，而不是通过修改已有代码来实现变化。
*/

// 抽象的银行业务员
type abstractBanker interface {
	//抽象的处理业务接口
	// 对外开放
	doBusiness()
}

// 存款的业务员
type saveBanker struct {
}

func (sb *saveBanker) doBusiness() {
	fmt.Println("进行了存款")
}

// 转账的业务员
type transferBanker struct {
	//AbstractBanker
}

func (tb *transferBanker) doBusiness() {
	fmt.Println("进行了转账")
}

type payBanker struct {
}

func (pb *payBanker) doBusiness() {
	fmt.Println("进行了支付")
}

/*******************    新业务 start  ****************************/
// 此时如果新业务，买股票，只有实现了abstractBanker接口即可
// 业务和之前业务完全解耦
type buyStockBanker struct {
}

func (bs *payBanker) buyStockBanker() {
	fmt.Println("买了股票")
}

/*******************    新业务 end  ****************************/

// 接口最大的意义就是多态思想
// 实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
// 多态
func bankerBusiness(banker abstractBanker) {
	//通过接口来向下调用，(多态现象)
	banker.doBusiness()
}

func main() {
	//进行存款
	sb := &saveBanker{}
	sb.doBusiness()

	//进行转账
	tb := &transferBanker{}
	tb.doBusiness()

	// 进行支付
	pb := &payBanker{}
	pb.doBusiness()

	// 多态
	bankerBusiness(&saveBanker{})
	bankerBusiness(&transferBanker{})
	bankerBusiness(&payBanker{})
}
