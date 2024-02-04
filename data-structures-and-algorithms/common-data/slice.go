package main

import "fmt"

/*
https://docs.golangjob.cn/advanced-go-programming-book/ch1-basic/ch1-03-array-string-and-slice.html
https://blog.csdn.net/fm_VAE/article/details/79266109

切片的截取
s[n]             切片s中索引位置为n的项
s[n:m]           从切片s的索引位置n到m-1处所获得的切片
s[n:]            从切片s的索引位置到len(s)-1处所获得的切片
s[:m]            从切片s的索引位置0到m-1处所获得的切片
s[:]             从切片s的索引位置0到len(s)-1处获得的切片

https://geektutu.com/post/hpg-slice.html
1. Copy
2. Append
3. Delete
4. Delete(GC)
5. Insert
6. Filter
7. Push
8. Pop
*/

/*
make([]type, length, capacity)
make([]type, length)

[]type{}
[]type{value1, value2,..., valueN}
*/
func createSlice() {
	// 字面量
	var s1 = []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("s1 addr:%p len:%d cap:%d s1:%v\n", s1, len(s1), cap(s1), s1)
	// s1 addr:0xc00000e4b0 len:6 cap:6 s1:[1 2 3 4 5 6]

	fmt.Println("\n")

	// make
	var s2 = make([]string, 10)
	s2 = append(s2, "hello")
	fmt.Printf("s2 addr:%p len:%d cap:%d s1:%v\n", s2, len(s2), cap(s2), s2)
	// s2 addr:0xc000072140 len:11 cap:20 s1:[          hello]

	fmt.Println("\n")
}

func testAppend() {
	list := make([]string, 0)
	list = append(list, "go", "c", "c++")
	fmt.Println(list) // [go c c++]
	println(list[1])  // c

	fmt.Println("-------------append------------------")
	list = append(list, "rust")                      // 追加1个元素
	list = append(list, []string{"vue", "react"}...) // 追加一个切片

	fmt.Println("list: ", list) // list:  [go c c++ rust vue react]

}

// 复制
func testCopy() {
	lang := []string{"go", "c", "c++"}
	newLang := make([]string, len(lang))

	// 深度复制
	copy(newLang, lang)

	fmt.Println("lang: ", lang)       // lang:  [go c c++]
	fmt.Println("newLang: ", newLang) // newLang:  [go c c++]

	newLang[0] = "vue"

	fmt.Println("lang: ", lang)       // lang:  [go c c++]
	fmt.Println("newLang: ", newLang) //  newLang:  [vue c c++]

	// 浅复制：赋值过程中发生的浅复制
	front := newLang
	fmt.Println("front: ", front) // front:  [vue c c++]

	// 浅复制 newLang切片中值也会跟着改变
	front[1] = "react"
	front[2] = "angular"
	fmt.Println("front: ", front)     // front:  [vue react angular]
	fmt.Println("newLang: ", newLang) // newLang:  [vue react angular]
}

/**
切片的截取
s[n]                            切片s中索引位置为n的项
s[n:m] 包含n, 不包含m             从切片s的索引位置n到m-1处所获得的切片
s[n:]  包含n，n后所有远素           从切片s的索引位置到len(s)-1处所获得的切片
s[:m]  不包含m, 0到m-1处所有元素，  从切片s的索引位置0到m-1处所获得的切片
s[:]   所有元素          		    从切片s的索引位置0到len(s)-1处获得的切片
*/

func cut() {
	fmt.Println("-----------------cut ---------------------")
	s := []string{"go", "c", "c++", "vue", "react", "rust"}

	// s[n]             切片s中索引位置为n的项
	fmt.Println("s[1] = ", s[1]) // s[1] =  c

	// s[n:m]  包含n, 不包含m     从切片s的索引位置n到m-1处所获得的切片
	s2 := s[1:3]
	fmt.Println("s2: ", s2) // s2:  [c c++]

	// s[n:] 包含n，n后所有远素    从切片s的索引位置到len(s)-1处所获得的切片
	s3 := s[1:]
	fmt.Println("s3:", s3)

	// s[:m] 不包含m, 0到m-1处所有元素，           从切片s的索引位置0到m-1处所获得的切片
	s4 := s[:3]
	fmt.Println("s4:", s4) // s4: s4: [go c c++]

	// s[:]  所有元素           从切片s的索引位置0到len(s)-1处获得的切片
	s5 := s[:]
	fmt.Println("s5:", s5) //s5: [go c c++ vue react rust]

}

func delSlice() {

	fmt.Println("----------删除操作----------")

	str := []string{"A", "B", "C", "D", "E", "F"}

	fmt.Println(str)

	//a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	//a = append(a[:i], a[i+N:]...) // 删除中间N个元素

	//a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	//a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	str = append(str[:3], str[4:]...) // 删除D元素
	fmt.Println("str: ", str)         // str:  [A B C E F]

}

type CityVO struct {
	Name      string    `json:"name"`
	Id        string    `json:"id"`
	LevelType string    `json:"levelType"`
	Children  []*CityVO `json:"children"`
}

func cityList() {
	hotCityList := make([]CityVO, 0)
	hotCityList = append(hotCityList,
		CityVO{
			Name:      "北京",
			Id:        "110100",
			LevelType: "2",
		},
		CityVO{
			Name:      "上海",
			Id:        "430100",
			LevelType: "2",
			Children:  make([]*CityVO, 0),
		},
		CityVO{
			Name:      "杭州",
			Id:        "330100",
			LevelType: "2",
		},
	)

	fmt.Println("hotCityList: ", hotCityList)
	// hotCityList:  [{北京 110100 2 []} {上海 430100 2 []} {杭州 330100 2 []}]

	fmt.Println(hotCityList[1].Name) // 上海

	fmt.Printf("%+v", hotCityList)
	// [{Name:北京 Id:110100 LevelType:2 Children:[]} {Name:上海 Id:430100 LevelType:2 Children:[]} {Name:杭州 Id:330100 LevelType:2 Children:[]}]

}

func main() {
	createSlice()

	testAppend()

	testCopy()

	cut()

	delSlice()

	cityList()
}
