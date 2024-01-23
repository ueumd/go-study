package main

import (
	"bytes"
	"fmt"
)

/**
https://blog.csdn.net/guolianggsta/article/details/125267881
*/

// 从Buffer中读取数据到指定容器
func testBuffer() {
	fmt.Println("testBuffer-----------------------------------\n")
	// c := make([]byte, 8)

	var b bytes.Buffer
	b.WriteString("hello")

	// b.Read(c) //一次读取8个byte到c容器中，每次读取新的8个byte覆盖c中原来的内容

	fmt.Println(b.String()) // hello

}

func testBuffer2() {
	fmt.Println("testBuffer2-----------------------------------\n")

	s := []byte("你好世界")
	fmt.Println(string(s))

	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	sep := []byte(",")

	fmt.Println(string(bytes.Join(s2, sep)))
}

func testReader() {
	fmt.Println("\ntestReader-----------------------------------")
	data := "123456789"

	re := bytes.NewReader([]byte(data))

	fmt.Println("re len", re.Len())
	fmt.Println("re size", re.Size())

	buf := make([]byte, 3)

	for {
		// 读取数据
		n, err := re.Read(buf) //一次读取3个byte到buf容器中，每次读取新的3个byte覆盖buf中原来的内容
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	testBuffer()
	testBuffer2()
	testReader()
}
