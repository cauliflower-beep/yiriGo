package main

import (
	"fmt"
	"time"
)

func child() {
	fmt.Println("kid:好啦，我开始写啦...")
	time.Sleep(time.Second * 5)
	//_panic("kid:这题我不会写QAQ")
	/*
		关于协程崩溃如何优雅重启，参考例子：
		https://blog.csdn.net/JusticeAngle/article/details/90614344
	*/
	fmt.Println("kid:我写完啦")
}

func father() {
	fmt.Println("dad:我看会儿电视，你去写作业！")
	go child()
	time.Sleep(time.Second * 3)
	fmt.Println("dad:没什么好节目...")
}
func mother() {
	fmt.Println("mom:我开始做饭啦，你写作业了没？")
	go child()
	time.Sleep(time.Second * 6)
	fmt.Println("mom:我做完饭啦，都来吃饭~")
}

func main() {
	// 1.父函数结束，子goroutine仍会继续执行
	//father()

	// 2.父goroutine结束，子goroutine仍会继续执行；
	//go father()

	// 3.子goroutine结束，父函数仍会继续执行
	//mother()

	// 4.子goroutine结束，父goroutine仍会继续执行
	//go mother()

	select {} // 死锁问题可见 ../issue/select.go 解释
}
