package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func main() {

	ww, err := exec.Command("CMD", "/C", " ping bing.com").Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(ww))

	ww, err = exec.Command("CMD", "/C", " ipconfig").Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	//	fmt.Println(string(ww))

	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	fmt.Printf("%q\n", reg.FindAllString(string(ww), -1)[0])

}
