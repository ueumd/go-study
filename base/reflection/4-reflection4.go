package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Id   int
	Name string
	Age  int
}

// 反射修改基本类型的值
func reflectChangeValue() {
	x := 123
	fmt.Println(x) //123

	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x) //999
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	// 修改结构体中的Name为BYEBYE
	f := v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("无效字段")
		return
	}

	//通过反射修改字段的值
	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}

func (p person) hello(name string) {
	fmt.Println("hello", name, ", my name is ", p.Name)
}

func main() {
	reflectChangeValue()

	p := person{1, "OK", 12}
	Set(&p)
	fmt.Println(p) //{1 OK 12}

	v := reflect.ValueOf(&p)
	mv := v.MethodByName("hello")

	args := []reflect.Value{reflect.ValueOf("joe")}

	mv.Call(args)
}
