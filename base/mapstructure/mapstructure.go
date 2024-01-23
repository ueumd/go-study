package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type blog struct {
	BlogId  string `mapstructure:"blogId"`
	Title   string `mapstructrue:"title"`
	Content string `mapstructure:"content"`
	Uid     string `mapstructure:"uid"`
	State   string `mapstructure:"state"`
}

type event struct {
	Type     string              `json:"type"`
	Database string              `json:"database"`
	Table    string              `json:"table"`
	Data     []map[string]string `json:"data"`
}

func testMapStructure() {
	e := event{}
	msg := []byte(`{ 
			"type": "UPDATE",  
			"database": "blog", 
			"table": "blog", 
			"data": [{"blogId": "100001",  "title": "title", "content": "this is a blog", "uid": "1000012", "state": "1" }]
		}`)

	if err := json.Unmarshal(msg, &e); err != nil {
		panic(err)
	}

	if e.Table == "blog" {
		var blogs []blog
		if err := mapstructure.Decode(e.Data, &blogs); err != nil {
			panic(err)
		}
		fmt.Println(blogs)
	}

}

func main() {
	testMapStructure()
}
