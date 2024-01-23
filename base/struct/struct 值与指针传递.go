package main

import "fmt"

/**
1. 结构体的方法只能和结构体在同一个包下
2. 内置的int类型不能加方法,int 在 main包下， 要加方法 可以重新定义一个int 类型
 */

type MyInt int

func (m MyInt) toString(price int)  {
	
}

type myCourse struct {
	name string
	price int
}

// 值传递
func (c myCourse) setPrice(price int)  {
	c.price = price
}

// 指针传递
func (this *myCourse) updatePrice(price int)  {
	this.price = price // (*this).price 引式解引用
}

func main()  {

	// 结构体是值传递
	cc := myCourse{"golang", 888}
	cc.setPrice(999)
	fmt.Println(cc.price) // 888

	cc.updatePrice(1000)
	fmt.Println(cc.price) // 1000
}