package main

import (
	"encoding/json"
	"fmt"
)

func map1() {
	// 声明
	var list map[int]string

	// 开辟内存
	list = make(map[int]string)

	list[0] = "hello"
	list[1] = "hello"
	list[2] = "hello"
	list[3] = "hello"

	fmt.Println(list)
}

type Node struct {
	Id       int     `json:"id"`
	ParentId int     `json:"parent_id"`
	Name     string  `json:"name"`
	Route    string  `json:"route"`
	Icon     string  `json:"icon"`
	Children []*Node `json:"children"`
}

func map2() {
	list := []*Node{
		{Id: 1, ParentId: 0, Name: "系统总览", Route: "/systemOverview", Icon: "icon-system"},
		//{Id: 2, ParentId: 0, Name: "系统配置", Route: "/systemConfig", Icon: "icon-config"},
		//{Id: 8, ParentId: 0, Name: "用户设置", Route: "/user", Icon: "icon-config"},

		{Id: 3, ParentId: 1, Name: "资产", Route: "/asset", Icon: "icon-asset"},
		{Id: 4, ParentId: 1, Name: "动环", Route: "/pe", Icon: "icon-pe"},

		//{Id: 5, ParentId: 2, Name: "菜单配置", Route: "/menuConfig", Icon: "icon-menu-config"},
		//{Id: 6, ParentId: 3, Name: "设备", Route: "/device", Icon: "icon-device"},
		//{Id: 7, ParentId: 3, Name: "机柜", Route: "/device", Icon: "icon-device"},
	}

	menuMap := make(map[int]*Node)

	for _, v := range list {
		if _, ok := menuMap[v.Id]; ok {
			v.Children = menuMap[v.Id].Children
			menuMap[v.Id] = v
		} else {
			// 给Children先开辟空间
			// c.Children = []
			v.Children = make([]*Node, 0)
			menuMap[v.Id] = v
		}

		if _, ok := menuMap[v.ParentId]; ok {
			menuMap[v.ParentId].Children = append(menuMap[v.ParentId].Children, menuMap[v.Id])
		} else {
			menuMap[v.ParentId] = &Node{Children: []*Node{menuMap[v.Id]}}
		}
	}

	fmt.Println(menuMap[0].Children)

	bytes, _ := json.MarshalIndent(menuMap[0].Children, "", "    ")
	fmt.Printf("%s\n", bytes)
}

func main() {
	map2()
}
