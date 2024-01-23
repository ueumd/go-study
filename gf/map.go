package main

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type User struct {
	Uid  int    `c:"uid"`
	Name string `c:"name"`
}

func main() {
	user := User{
		Uid:  1,
		Name: "john",
	}

	// 结构体转Map
	userMap := gconv.Map(user)
	userMap1 := gconv.Map(&user)
	g.Dump(userMap)
	g.Dump(userMap1)

	println(user.Name)

	newUser := User{}

	// Map -> Struct
	if err := gconv.Struct(userMap, &newUser); err == nil {
		g.Dump(newUser)
		println(newUser.Name)
	} else {
		g.Dump(err)
	}

	// 序列化 反序列化
	//序列化 → []byte
	j, _ := json.Marshal(user)
	j1, _ := json.Marshal(userMap)

	println(j)
	println(j1)

	//反序列化 → struct
	user2 := User{}
	json.Unmarshal(j, &user2)

	user3 := User{}
	json.Unmarshal(j1, &user3)

	println(user3.Name)

}
