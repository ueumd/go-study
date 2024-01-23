package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type myMath struct {
	a int
}

func (mm *myMath) Add(num1 float64, num2 float64) float64 {
	reply := num1 + num2
	return reply
}

func (mm *myMath) Sub(num1 float64, num2 float64) float64 {
	reply := num1 - num2
	return reply
}

func main() {
	myMathStruct := new(myMath)

	myMathStructType := reflect.TypeOf(myMathStruct) //  返回obj的对象类型
	myMathStructValue := reflect.ValueOf(myMathStruct)

	fmt.Println(myMathStructType) // *main.myMath

	fmt.Printf("TypeOf: %#v \n", myMathStructType)   //  &reflect.rtype{t:abi.Type{Size_:0x8, PtrBytes:0x8, Hash:0xb6cf8ea5, TFlag:0x9, Align_:0x8, FieldAlign_:0x8, Kind_:0x36, Equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0xa46e20), GCData:(*uint8)(0xb62008), Str:10653, PtrToThis:0}}
	fmt.Printf("valueOf: %#v \n", myMathStructValue) // &main.myMath{a:0}

	//遍历方法
	for i := 0; i < myMathStructType.NumMethod(); i++ {
		method := myMathStructType.Method(i)
		fmt.Println("method:" + method.Name) // 打印方法名 Add Sub

		fun := reflect.ValueOf(myMathStruct).MethodByName(method.Name)
		methodType := method.Type
		args := make([]reflect.Value, methodType.NumIn()-1)

		//遍历参数
		for j := 1; j < methodType.NumIn(); j++ {
			//参数类型
			arg_type := methodType.In(j).Kind().String()

			fmt.Println("参数" + strconv.Itoa(j) + ":" + arg_type)
		}
		ret := fun.Call(args)
		fmt.Println(ret[0])

		/**
		method:Add
		参数1:float64
		参数2:float64
		3
		method:Sub
		参数1:float64
		参数2:float64
		-1
		*/
	}

}
