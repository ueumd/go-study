package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"userName"`
	Age  int32  `json:"age"`
	sex  string `json:"sex"`
}

func init() {
	var user User
	user.Name = "小明"
	user.Age = 18
	user.sex = "男"

	res, _ := json.Marshal(user)
	fmt.Println(string(res)) //{"userName":"小明","age":18}
}

func main() {

}
