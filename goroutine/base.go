package main

import (
	"fmt"
	"runtime"
)

/*
golang中的协程可以理解为用户级线程，它是对内核透明的。
也就是系统并不知道有协程的存在，是完全由用户自己的程序调度的。

golang一大特色就是从语言层面原生支持协程，在函数或者方法前加 go 关键字就可以创建一个协程。可以说golang中的协程就是goroutine.

go运行时的调度器使用 runtime.GOMAXPROCS() 函数来确定需要使用多少个 OS 线程来同时执行go代码。
Go1.5版本之前，默认使用的是单核心执行，之后默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把go代码同时调度到8个os线程上跑。

*/

func main() {
	// 获取当前机器上的cpu核心数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum : ", cpuNum)

	// 设置使用多个cpu执行goroutine
	runtime.GOMAXPROCS(16)
	fmt.Println("ok")
}
