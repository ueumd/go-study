package main

import "fmt"

const OK = 0

var msgCode = map[int]string{
	OK: "Hello",
}

func main() {
	fmt.Println(msgCode[OK])
}
