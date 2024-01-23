package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func init() {
	type H map[string]interface{}

	pro := H{
		"name":  "安徽省",
		"items": [...]string{"合肥", "黄山"},
	}
	fmt.Println(pro)
}

func main() {
	var provinceList []interface{}
	var districtList []interface{}
	var province map[string]interface{}
	var m map[string]interface{}

	m = make(map[string]interface{})
	m["name"] = "准南"
	m["items"] = [...]string{"田家俺", "寿县"}
	districtList = append(districtList, m)

	m = make(map[string]interface{})
	m["name"] = "合肥"
	m["items"] = [...]string{"肥东", "肥西"}
	districtList = append(districtList, m)

	province = make(map[string]interface{})
	province["name"] = "安徽"
	province["items"] = districtList

	provinceList = append(provinceList, province)

	jsonStr, err := json.Marshal(provinceList)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s\n", jsonStr)
}
