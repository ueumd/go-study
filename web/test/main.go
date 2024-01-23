package main

import (
	"fmt"
	"strings"
)

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

func main() {
	// parsePattern("/hello/:id")
	// ("/assets/*filepath")

	str := make([]string, 0)

	//str[0] = "hello"
	//str[1] = "world"

	item := "*world"

	// 动态扩容
	str = append(str, "hello")
	str = append(str, item)

	fmt.Println(str) // [hello world]

	fmt.Println(item[0] == '*')

}
