package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
Web是基于http协议的一个服务。
Go语言提供了一个完善的net/http包，可以很方便的搭建起Web服务。同时能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作。

Go不需要PHP起web服务时候的nginx、apache服务器这些，他直接就监听tcp端口了，做了nginx做的事情。
*/

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello 广志!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayHelloName)       //设置访问的路由
	err := http.ListenAndServe(":8989", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
