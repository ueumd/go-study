package main

import "fmt"

/*============================ 耦合度极高的设计 ========================================*/
// === > 奔驰汽车 <===
type benzCar struct {
}

func (this *benzCar) run() {
	fmt.Println("Benz is running...")
}

type bmwCar struct {
}

func (this *bmwCar) run() {
	fmt.Println("BMW is running...")
}

// ===> 司机张三  <===
type zhang3 struct{}

func (this *zhang3) driveBenz(benz *benzCar) {
	fmt.Println("zhang3 Drive Benz")
	benz.run()
}

func (this *zhang3) driveBmw(bmw *bmwCar) {
	fmt.Println("zhang3 Drive BMW")
	bmw.run()
}

// ===> 司机李四 <===
type li4 struct{}

func (this *li4) driveBenz(benz *benzCar) {
	fmt.Println("li4 Drive Benz")
	benz.run()
}

func (this *li4) driveBmw(bmw *bmwCar) {
	fmt.Println("li4 Drive BMW")
	bmw.run()
}

/*
耦合度太高
再添加 汽车 或 司机 都要实现，所有的模块耦合度极高
*/

func main_() {
	/*
		小规模设计尚可，但是一旦要扩展，比如我现在要增加一个丰田汽车 或者 司机王五
		那么模块和模块的依赖关系将成指数级递增，想蜘蛛网一样越来越难维护和捋顺。
	*/

	//业务1 张3开奔驰
	benz := &benzCar{}
	zs := &zhang3{}
	zs.driveBenz(benz)

	//业务2 李四开宝马
	bmw := &bmwCar{}
	ls := &li4{}
	ls.driveBmw(bmw)
}

/*

=================================================================================

利用接口，将模块分为3个层次： 抽象层、实现层、业务逻辑层

依赖倒转原则：
1. 先将抽象层的模块和接口定义出来，这里里就需要了interface接口的设计
2. 然后我们依照抽象层，依次实现每个实现层的模块，在我们写实现层代码的时候，
实际上我们只需要参考对应的抽象层实现就好了，实现每个模块，也和其他的实现的模块没有关系，这样也符合了上面介绍的开闭原则。

3. 这样实现起来每个模块只依赖对象的接口，而和其他模块没关系，依赖关系单一。系统容易扩展和维护。

4. 我们在指定业务逻辑也是一样，只需要参考抽象层的接口来业务就好了，
抽象层暴露出来的接口就是我们业务层可以使用的方法，然后可以通过多态的线下，接口指针指向哪个实现模块，调用了就是具体的实现方法，
这样我们业务逻辑层也是依赖抽象成编程。
*/

// ===== >   抽象层  < ========

// 抽像汽车
type car interface {
	run()
}

// 抽像司机
type Driver interface {
	drive(car2 car)
}

// ===== >   实现层  < ========
type benzCar2 struct{}

func (this *benzCar2) run() {
	fmt.Println("Benz is running...")
}

type bmwCar2 struct{}

func (this *bmwCar2) run() {
	fmt.Println("BMW is running...")
}

// 司机
type zhang_3 struct{}

func (this *zhang_3) drive(car2 car) {
	fmt.Println("Zhang3 drive car")
	car2.run()
}

type li_4 struct{}

func (this *li_4) drive(car2 car) {
	fmt.Println("li4 drive car")
	car2.run()
}

// ===== >   业务逻辑层  < ========
func main() {

	//张3 开 宝马
	var bmw car
	bmw = &bmwCar2{}

	var zhang3 Driver
	zhang3 = &zhang_3{}

	zhang3.drive(bmw)

	//李4 开 奔驰
	var benz car
	var li4 Driver

	benz = &bmwCar2{}
	li4 = &li_4{}
	li4.drive(benz)
}
