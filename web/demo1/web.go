package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func SayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func Login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //获取请求的方法
	if req.Method == "GET" {
		t, _ := template.ParseFiles("front/login1.gtpl")
		t.Execute(res, nil)
	} else {
		req.ParseForm()

		//req.Form是一个url.Values类型，里面存储的是对应的类似key=value的信息
		v := url.Values{}
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		v.Add("friend", "Zoe")

		fmt.Println(v.Get("name"))
		fmt.Println(v.Get("friend")) // 打印出: Jess
		fmt.Println(v["friend"])     // 打印出: [Jess Sarah Zoe]

		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", req.Form["username"])
		fmt.Println("password:", req.Form["password"])
	}
}

func Login2(res http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //获取请求的方法
	if req.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("front/login2.gtpl")
		t.Execute(res, token)

	} else {

		//请求的是登陆数据，那么执行登陆的逻辑判断
		req.ParseForm()
		token := req.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		fmt.Println("username length:", len(req.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(req.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(req.Form.Get("password")))
		template.HTMLEscape(res, []byte(req.Form.Get("username"))) //输出到客户端
	}
}

func init() {
	http.HandleFunc("/", SayhelloName) //设置访问的路由
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login2", Login2)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
