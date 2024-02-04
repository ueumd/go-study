package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/**
https://www.zhihu.com/question/400373520

什么时候用结构体指针
简而言之，以下情况推荐使用结构体指针：

1. 需要改变原来结构体的值
2. 为了节约传递的成本, 结构体很大，传指针避免复制
3. 需要对结构体进行修改
4. 结构体对象复制会导致失效，因此不能复制，比如锁
5. 参数可能是空值，传指针便于判断是否为nil
*/

/**

https://golang.dbwu.tech/performance/struct_slice/

业务开发中，一个常见的场景是将多个相同类型的 结构体 变量存入一个数据容器中，通常我们会使用 切片 作为数据容器。
那么对于结构体来说，存储其值和存储其指针，性能差异有多大呢？
performance2_test.go
performance_test.go
*/

type TagInfo struct {
	Tag     string `form:"tag" json:"tag"`
	Version int    `form:"version" json:"version"`
}

type TagList struct {
	// 指针数组 存地址
	Tags []*TagInfo `form:"tags" json:"tags"`
}

func getTags() *[]TagInfo {
	tags := make([]TagInfo, 0)

	for i := 1; i <= 10; i++ {
		tags = append(tags, TagInfo{
			Tag:     "tag-" + strconv.Itoa(i),
			Version: i,
		})
	}
	return &tags
}

func main() {

	tags := getTags()
	fmt.Println("tags: ", tags)
	// tags:  &[{tag-1 1} {tag-2 2} {tag-3 3} {tag-4 4} {tag-5 5} {tag-6 6} {tag-7 7} {tag-8 8} {tag-9 9} {tag-10 10}]

	res := TagList{
		Tags: make([]*TagInfo, 0),
	}

	for _, tag := range *tags {
		res.Tags = append(res.Tags, &tag)
	}

	fmt.Println("res: ", res)
	// res:  {[0xc000094168 0xc000094180 0xc000094198 0xc0000941b0 0xc0000941c8 0xc0000941e0 0xc0000941f8 0xc000094210 0xc000094228 0xc000094240]}
	fmt.Println(res.Tags[2].Tag, res.Tags[2].Version) // tag-3 3

	bytes, _ := json.MarshalIndent(res, "", "    ")
	fmt.Printf("%s\n", bytes)

}
