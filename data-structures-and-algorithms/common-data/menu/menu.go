package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Id       int     `json:"id"`
	ParentId int     `json:"parent_id"`
	Name     string  `json:"name"`
	Route    string  `json:"route"`
	Icon     string  `json:"icon"`
	Children []*Node `json:"children"`
}

func getTreeIterative(list []*Node, parentId int) []*Node {
	memo := make(map[int]*Node)
	for _, v := range list {
		if _, ok := memo[v.Id]; ok {
			v.Children = memo[v.Id].Children
			memo[v.Id] = v
		} else {
			v.Children = make([]*Node, 0)
			memo[v.Id] = v
		}
		if _, ok := memo[v.ParentId]; ok {
			memo[v.ParentId].Children = append(memo[v.ParentId].Children, memo[v.Id])
		} else {
			memo[v.ParentId] = &Node{Children: []*Node{memo[v.Id]}}
		}
	}
	return memo[parentId].Children

}

func main() {
	list := []*Node{
		{Id: 1, ParentId: 0, Name: "系统总览", Route: "/systemOverview", Icon: "icon-system"},
		{Id: 2, ParentId: 0, Name: "系统配置", Route: "/systemConfig", Icon: "icon-config"},
		{Id: 8, ParentId: 0, Name: "用户设置", Route: "/user", Icon: "icon-config"},

		{Id: 3, ParentId: 1, Name: "资产", Route: "/asset", Icon: "icon-asset"},
		{Id: 4, ParentId: 1, Name: "动环", Route: "/pe", Icon: "icon-pe"},

		{Id: 5, ParentId: 2, Name: "菜单配置", Route: "/menuConfig", Icon: "icon-menu-config"},
		{Id: 6, ParentId: 3, Name: "设备", Route: "/device", Icon: "icon-device"},
		{Id: 7, ParentId: 3, Name: "机柜", Route: "/device", Icon: "icon-device"},
	}

	for index, v := range list {
		fmt.Println(index, v)
	}

	res := getTreeIterative(list, 0)
	fmt.Println(res)

	bytes, _ := json.MarshalIndent(res, "", "    ")
	fmt.Printf("%s\n", bytes)
}
