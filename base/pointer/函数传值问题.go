package main

import (
	"fmt"
	"time"
)

type Tt struct {
	Name     string
	Age      int
	Email    string
	Addr     string
	Nickname string
}

func Abc1() {
	fmt.Println("start")
	starttime := time.Now()
	var data Tt
	for i := 0; i < 10000; i++ {
		Set1(data)
	}
	fmt.Println("starttime:", time.Since(starttime))

}

func Set1(data Tt) Tt {
	data.Name = "ashan"
	data.Age = 32
	data.Email = "abc@abc.com"
	data.Addr = "china"
	data.Nickname = "ashan"
	return data
}

func Abc() {
	fmt.Println("start")
	starttime := time.Now()
	var data Tt
	for i := 0; i < 10000; i++ {
		Set(&data)
	}
	fmt.Println(time.Since(starttime))

}

func Set(data *Tt) *Tt {
	data.Name = "ashan"
	data.Age = 32
	data.Email = "abc@abc.com"
	data.Addr = "china"
	data.Nickname = "ashan"
	return data
}

func Setint(data *int) *int {
	*data = 5 + 12*8
	return data
}
