package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	案例比较互斥锁和原子操作的性能
*/

var (
	x  int64
	mx sync.Mutex
	wg sync.WaitGroup
)

// 普通函数，并发不安全
func Add() {
	x++
	wg.Done()
}

// 互斥锁，并发安全，性能低于原子操作
func MxAdd() {
	mx.Lock()
	x++
	mx.Unlock()
	wg.Done()
}

// 原子操作，并发安全，性能高于互斥锁，只针对go中的一些基本数据类型使用
func AmAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {

	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go Add() // 普通版Add函数不是并发安全的
		//go MxAdd() // 加锁版Add函数，是并发安全的，但是加锁性能开销大
		go AmAdd() // 原子操作版Add函数，是并发安全的，性能优于加锁版
	}

	end := time.Now()
	wg.Wait()
	fmt.Println(x)
	fmt.Println(end.Sub(start)) // 开始到结束的持续时间

}
