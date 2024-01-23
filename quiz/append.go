package main

import "fmt"

func init() {
	s := []int{1, 2, 3}
	ss := s[1:]
	for i := range ss {
		ss[i] += 10
	}
	fmt.Println(s)
	fmt.Println(ss)
	ss = append(ss, 4) //这里不再是引用了
	for i := range ss {
		ss[i] += 10
	}
	fmt.Println(s)
	fmt.Println(ss)
}
