package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

/*
golang中的协程可以理解为用户级线程，它是对内核透明的。
也就是系统并不知道有协程的存在，是完全由用户自己的程序调度的，依赖于go语言运行时自身提供的调度器。

golang一大特色就是从语言层面原生支持协程，在函数或者方法前加 go 关键字就可以创建一个协程。可以说golang中的协程就是goroutine.

go运行时的调度器使用 runtime.GOMAXPROCS() 函数来确定需要使用多少个 OS 线程来同时执行go代码。
Go1.5版本之前，默认使用的是单核心执行，之后默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把go代码同时调度到8个os线程上跑。

*/

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
	// 但其实对于IO密集型的场景，可以把GOMAXPROCS的值超过CPU核数，在笔者维护的某个服务中，将GOMAXPROCS设为CPU核数的2倍，压测结果表明，吞吐能力大概能提升10%
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
	fmt.Println("hello goroutine~")
	time.Sleep(3 * time.Second) // 等待子goroutine执行完毕
}
