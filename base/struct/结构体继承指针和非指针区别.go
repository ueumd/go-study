package main

import "fmt"

type iter interface {
	run()
	sleep()
}

type animalStruct struct {
	name string
}

func (a *animalStruct) run() {
	fmt.Println(a.name, " ::run()")
}

func (a animalStruct) sleep() {
	fmt.Println(a.name, " ::sleep()")
}

type dogStruct struct {
	animalStruct // 字段继承
}

type catStruct struct {
	*animalStruct // 字段指针继承
}

func ts1() {
	// 实例[字段继承]
	dog := dogStruct{animalStruct{"张铁柱"}}
	dog.run()   // 张铁柱  ::run()
	dog.sleep() // 张铁柱  ::sleep()

	// 指针实例[字段继承]
	dog2 := &dogStruct{animalStruct{"旺财"}}
	dog2.run()   // 旺财  ::run()
	dog2.sleep() // 旺财  ::sleep()

	// 实例[字段指针继承]
	cat := catStruct{&animalStruct{"小花"}}
	cat.run()   // 	小花  ::run()
	cat.sleep() // 小花  ::sleep()

	// 指针实例[字段指针继承]
	cat2 := &catStruct{&animalStruct{"汤姆"}}
	cat2.run()   // 	汤姆  ::run()
	cat2.sleep() // 汤姆  ::sleep()
}

func main() {

	// ======实例转换为接口=============
	var it iter

	// 实例[字段继承]
	// dog := dogStruct{animalStruct{"张铁柱"}}

	// cannot use dog (variable of type dogStruct) as iter value in assignment:
	// dogStruct does not implement iter (method run has pointer receiver)
	//it = dog // 错误  run方法是指针接收者 dog和animalStruct 是都不是指针实例， 父类没有实现 run 方式

	//it.run()   //  父类没有实现 run 方式
	//it.sleep() // 张铁柱  ::sleep()

	// 指针实例[字段继承]
	dog2 := &dogStruct{animalStruct{"旺财"}}
	it = dog2
	it.run()   // 旺财  ::run()
	it.sleep() // 旺财  ::sleep()

	// 实例[字段指针继承]
	cat := catStruct{&animalStruct{"小花"}}
	it = cat
	it.run()   // 	小花  ::run()
	it.sleep() // 小花  ::sleep()

	// 指针实例[字段指针继承]
	cat2 := &catStruct{&animalStruct{"汤姆"}}
	it = cat2
	it.run()   // 	汤姆  ::run()
	it.sleep() // 汤姆  ::sleep()

}
