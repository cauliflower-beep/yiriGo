package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func goA(ch chan int) {
	fmt.Println("A===")
	ch <- 1
	wg.Done()
}

func goB(ch chan int) {
	<-ch
	fmt.Println("B===")
	wg.Done()
}

func main() {
	wg.Add(6)
	// 使用chan控制同步，交替打印 A/B 共三次
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go goA(ch)
		go goB(ch)
	}
	wg.Wait()
}
