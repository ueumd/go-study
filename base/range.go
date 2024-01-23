package main

import "fmt"

func main() {
	x := []string{"a", "b", "C"}

	for index := range x {
		fmt.Println(index)
	}

	fmt.Println("---------------------------\n")

	for index, value := range x {
		fmt.Println(index, value)
	}

	/**
	0
	1
	2
	---------------------------

	0 a
	1 b
	2 C
	*/
}
