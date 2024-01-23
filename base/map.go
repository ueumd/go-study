package main

import "fmt"

func Testmap() {
	rating := map[string]float32{"c": 5, "go": 4.5, "c++": 11}
	csharpRating, ok := rating["c#"]

	if ok {
		fmt.Println("%s is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
}
