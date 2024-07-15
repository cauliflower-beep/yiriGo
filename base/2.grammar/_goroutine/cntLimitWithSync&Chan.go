package main

import (
	"fmt"
	"math"
	"runtime"
)

//var wg = sync.WaitGroup{}

func process(ch chan struct{}, i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())

	<-ch

	wg.Done()
}

func main() {
	// 模拟用户需求go业务的数量
	taskCnt := math.MaxInt64

	ch := make(chan struct{}, 3)

	for i := 0; i < taskCnt; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go process(ch, i)
	}

	wg.Wait()
}
