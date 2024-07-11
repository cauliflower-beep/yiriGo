package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
goroutine的大小并不是一成不变的。
每个goroutine需要能够独立运行，所以他们要有自己独立的栈。
假如每个goroutine分配固定的栈大小并且不能增长，太小则会导致溢出，太大又会浪费空间，无法存在许多的goroutine。
为了解决这个问题，goroutine初始时只给栈分配很小的空间，然后随着使用过程自动的增长。这也是为什么Go可以开千千万万个goroutine而不会耗尽内存。

实现方式:
每次执行函数调用时，go的runtime都会进行检测。若当前栈的大小不够用，则会触发“终中断”，从当前函数进入到go的运行时库。go的运行时库会保存此时
的函数上下文环境，然后分配一个新的足够大的栈空间，将旧栈的内容拷贝到新栈中，并做一些设置，使得当函数恢复运行时，函数会在新分配的栈中继续执行，
仿佛整个过程都没发生过一样。函数视角来看，会觉得自己使用的是一块儿“无限大小”的栈空间。
*/
const numGoroutines = 1e4

/*
计算一个 _goroutine 有多大
*/
func main() {
	memConsumed := func() uint64 {
		runtime.GC()           // gc,排除对象影响
		var s runtime.MemStats // 记录内存统计信息
		runtime.ReadMemStats(&s)
		return s.Sys // 从操作系统获得的内存总字节数
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c // 防止goroutine退出，内存被释放
	}
	wg.Add(numGoroutines)
	before := memConsumed() // goroutine创建前的内存
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() // goroutine创建后的内存

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
	fmt.Println("")
}
