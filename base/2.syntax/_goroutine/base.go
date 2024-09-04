package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkLink(link string) {
	runtime.Gosched() // 将时间片让给别人，下次某个时候继续恢复执行该goroutine
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

func main() {
	// 获取当前机器上的cpu核心数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum : ", cpuNum)

	// 告诉调度器同时使用多个线程执行goroutine，并返回之前的设置。如果n<1,不会改变当前设置
	// 但其实对于IO密集型的场景，可以把GOMAXPROCS的值超过CPU核数。例如在某些服务中，将GOMAXPROCS设为CPU核数的2倍，压测结果表明，吞吐能力大概能提升10%
	runtime.GOMAXPROCS(16)
	fmt.Println("ok")

	links := []string{
		"https://book.douban.com/",
		"https://translate.google.cn/",
		"https://www.bilibili.com/",
		"https://github.com/",
		"https://www.baidu.com/",
	}
	/*
		线性程序:有很严重的性能问题：
		必须等待前一个请求执行完毕，后一个请求才能继续执行。如果有一个请求的网站出现了问题，则可能需要等待很长时间。
		这种情况在网络访问、磁盘文件访问时经常会遇到。
	*/
	//for _, link := range links {
	//	checkLink(link)
	//}

	// 并发执行
	for _, link := range links {
		go checkLink(link)
	}
	fmt.Println("hello _goroutine~")
	time.Sleep(3 * time.Second) // 等待子goroutine执行完毕
}
