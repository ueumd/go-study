package main

import (
	"fmt"
)

type Response[T any] struct {
	Code    int
	Message string
	Data    any
}

func testAny() {
	res := Response[string]{
		Code:    0,
		Message: "OK",
		Data:    1,
	}

	fmt.Println("res: ", res) // res:  {0 OK 1}
}

func testAny2() {
	lang := make(map[any]string, 0)
	lang[0] = "vue"
	lang[1] = "react"
	lang["name"] = "Tom"
	lang["age"] = "1"

	fmt.Println(lang["name"]) // Tom

	if age, ok := lang["age"]; ok {
		fmt.Println("age: ", age) // age:  1
	}

	fmt.Println(len(lang)) // 4

	for k, v := range lang {
		println("k:", k, "v:", v)
		/**
		k: (0xccc880,0xd00038) v: vue
		k: (0xccc880,0xd00030) v: react
		k: (0xccc700,0xd00180) v: Tom
		k: (0xccc700,0xd00190) v: 1
		*/
	}

	delete(lang, 1)   // 删除react
	fmt.Println(lang) // map[age:1 name:Tom 0:vue]
}

func main() {
	testAny()
	testAny2()

}
