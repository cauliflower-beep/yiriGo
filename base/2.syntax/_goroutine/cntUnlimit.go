package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	// 模拟用户需求业务的数量
	taskCnt := math.MaxInt64

	for i := 0; i < taskCnt; i++ {
		go func(i int) {
			// do sth
			fmt.Println("go func ", i, " goroutine cnt = ", runtime.NumGoroutine())
		}(i)
	}
}
