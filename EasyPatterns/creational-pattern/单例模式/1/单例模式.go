package main

import "fmt"

/*
懒汉式: 首次使用时创建单例实例
*/

type singleton struct {
}

func (this *singleton) connection() {
	println("connection method")
}

var instance *singleton

func getInstance() *singleton {

	//只有首次GetInstance()方法被调用，才会生成这个单例的实例
	if instance == nil {
		instance = new(singleton)
		fmt.Println("-------------created instance----------------")
	}

	// 已有直接返回
	return instance
}

func main() {

	// 单线程
	//sql := getInstance()
	//sql.connection()

	for i := 0; i <= 100; i++ {
		go func(i int32) {
			println("i: ", i)
			sql := getInstance()
			sql = getInstance()
			sql = getInstance()
			sql = getInstance()
			sql = getInstance()
			sql.connection()
		}(int32(i))
	}

	for {
	}

}
