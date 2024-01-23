package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type user struct {
	//User字段
	Id   int
	Name string
	Age  int
}

// User方法
func (u user) Hello(name string, arg2 int, arg3 float64) {
	fmt.Println("hi " + name)
}

func (u user) Work(who string, arg2 int) {
	fmt.Println(who + " is busy...")
}

func (u user) TestCallMethod(name string) string {
	return name
}

func (u user) secret(salary int) {
	fmt.Println("secret...")
}

// 反射的基本操作
func testReflect(info interface{}) {
	//取类型
	infoType := reflect.TypeOf(info)

	//取值
	infoValue := reflect.ValueOf(info)

	fmt.Println("Type: ", infoType.Name(), "Kind: ", infoType.Kind()) // Type:  user Kind:  struc

	if k := infoType.Kind(); k != reflect.Struct {
		fmt.Println("类型错误")
		return
	}

	// 打印字段
	for i := 0; i < infoType.NumField(); i++ {
		field := infoType.Field(i)                   //取到字段
		fieldValue := infoValue.Field(i).Interface() //取到字段对应的值
		fmt.Printf("%6s: %v = %v \n", field.Name, field.Type, fieldValue)
		/**
		  Id: int = 1
		  Name: string = ok
		   Age: int = 12
		*/
	}

	//NumMethod() 方法集: 该类型的方法集，Type类型提供了方法来返回方法数量，访问各个方法。reflect包定义了Method类型来表示一个方法

	// 打印方法 (私有方法不会打印)
	for i := 0; i < infoType.NumMethod(); i++ {
		method := infoType.Method(i)
		fmt.Printf("%6s: %v\n", method.Name, method.Type)
		//  Hello: func(main.user, string)
		//  Work: func(main.user, string)

		methodType := method.Type

		/**
		*函数类型
		func IsVariadic() bool    参数是否可变
		func NumIn() int          参数的数量，需要注意的是，可变参数单独作为slice计算
		func NumOut() int         返回值的数量，需要注意的是，可变参数单独作为slice计算
		func In(i int) Type       第i个参数，i从0开始
		func Out(i int) Type      第i个返回值，i从0开始
		*/
		for j := 1; j < methodType.NumIn(); j++ {
			argType := methodType.In(j).Kind().String()
			fmt.Println(method.Name + " 参数" + strconv.Itoa(j) + ": " + argType)
		}

	}
}

// 调用结构体方法
func callStructMethod(v *reflect.Value, method string, params []interface{}) {
	// 字符串方法调用，且能找到实例conf的属性.Op
	f := (*v).MethodByName(method)
	if f.IsValid() {
		args := make([]reflect.Value, len(params))
		for k, param := range params {
			args[k] = reflect.ValueOf(param)
		}
		// 调用
		ret := f.Call(args)
		if ret[0].Kind() == reflect.String {
			fmt.Printf("%s Called result: %s\n", method, ret[0].String())
		}
	} else {
		fmt.Println("can't call " + method)
	}
	fmt.Println("")
}

func main() {
	u := user{1, "ok", 12}
	// testReflect(u)

	uu := reflect.ValueOf(u)
	callStructMethod(&uu, "TestCallMethod", []interface{}{"Tom"})
	// TestCallMethod Called result: Tom

	// testReflect(100) //类型错误
}
