package main

import (
	"fmt"
	"runtime"
	"sync"
)

const numGoroutines = 1e4
/*
计算一个 goroutine 有多大
 */
func main(){
	memConsumed := func()uint64 {
		runtime.GC()		// gc,排除对象影响
		var s runtime.MemStats 	// 记录内存统计信息
		runtime.ReadMemStats(&s)
		return s.Sys 	// 从操作系统获得的内存总字节数
	}

	var c <- chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c	// 防止goroutine退出，内存被释放
	}
	wg.Add(numGoroutines)
	before := memConsumed()		// goroutine创建前的内存
	for i := numGoroutines;i >0 ;i--{
		go noop()
	}
	wg.Wait()
	after := memConsumed()		// goroutine创建后的内存

	fmt.Printf("%.3fkb",float64(after-before)/numGoroutines/1000)
	fmt.Println("")
}