package main

import (
	"fmt"
	"time"
)

/*
main函数结束的时候，所有goroutine都会跟着一起结束。这个是老生常谈。否则我们也不会在各种goroutine初始教程里，看到time.sleep
这种让 main 等待goroutine执行完成的操作。

但main函数未结束时，父goroutine以及由它发起的子goroutine的生命周期就是另外一套逻辑了：
	1.父函数结束，子goroutine仍会继续执行；
	2.父goroutine结束，子goroutine仍会继续执行；
*/

func child() {
	fmt.Println("kid:好啦，我开始写啦...")
	time.Sleep(time.Second * 5)
	//panic("kid:这题我不会写QAQ")
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
	fmt.Println("mom我做完饭啦，都来吃饭~")
}

/*
上文说到:
main函数退出，所有协程退出；
父子协程相互独立，若父协程退出，不影响子协程的运行。
项目中要格外注意父子协程的退出关系，以免造成大量的孤儿协程。

问题来了，如果我们想要优雅的关闭子协程，该怎么操作呢?
--答案就是借助context了.可以参考 ../_context 中的案例加以理解。
当然也可以通过channel来解决，具体自己练习。
*/

func main() {
	// 1.父函数结束，子goroutine仍会继续执行
	//father()

	// 2.父goroutine结束，子goroutine仍会继续执行；
	//go father()

	// 3.子goroutine结束，父函数仍会继续执行
	mother()

	// 4.子goroutine结束，父goroutine仍会继续执行
	//go mother()

	select {} // 死锁问题可见 ../issue/select.go 解释
}
