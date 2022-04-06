package main

import (
	"fmt"
	"time"
)



func addNum(){
	for i := 0;i <100; i++{
		lock.Lock()
		num += 1
		lock.Unlock()
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	start := time.Now()
	go addNum()
	go addNum()
	wg.Wait()
	fmt.Println(num)	// 结果永远为200
	fmt.Println("运行时间：",time.Since(start))
}
