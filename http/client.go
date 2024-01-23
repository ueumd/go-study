package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
JSONPlaceholder 使用方式非常简单，提供了 GET、POST、PUT、PATCH、DELETE 几个请求方法。
http://jsonplaceholder.typicode.com/
*/

func main1() {
	res, err := http.Get("https://www.baidu.com")

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(res.Body)

	// 关闭连接
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data) // 输出 HTML代码

}

func main2() {
	client := &http.Client{}
	url := "http://jsonplaceholder.typicode.com/posts"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Charset", "utf-8")
	//req.Header.Add("Accept-Encoding","br, gzip, deflate")
	req.Header.Add("Accept-Language", "zh-cn")
	req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Cookie","xxxxxxxxxxxxxxx")
	//req.Header.Add("Content-Lenght",xxx)
	req.Header.Add("Host", "www.baidu.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

	rep, err := client.Do(req) //发起请求

	data, err := io.ReadAll(rep.Body)

	// 关闭连接
	rep.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data) // 输出 HTML代码
}

/*
*
POST
*/
func main() {
	client := &http.Client{}

	req_data := `{"userId":1 , "title":"app", "body": "Hello App"}`
	url := "http://jsonplaceholder.typicode.com/posts"

	req, err := http.NewRequest("POST", url, strings.NewReader(req_data))
	if err != nil {
		log.Fatal(err)
	}
	//Content-Type很重要，下文解释
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "multipart/form-data")

	rep, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(rep.Body)

	rep.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)
	/**
	{
	  "userId": 1,
	  "title": "app",
	  "body": "Hello App",
	  "id": 101
	}
	*/
}
