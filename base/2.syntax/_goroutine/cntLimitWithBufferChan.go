package main

import (
	"fmt"
	"math"
	"runtime"
)

func handler(ch chan bool, i int) {
	// do sth.
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch // 无值可读时会阻塞
}

func main() {
	// 模拟用户需求业务的数量
	taskCnt := math.MaxInt64

	// taskCnt := 10

	ch := make(chan bool, 3)

	for i := 0; i < taskCnt; i++ {
		ch <- true // 通道缓冲区填满时会阻塞，直到有空间可用
		go handler(ch, i)
	}
}
