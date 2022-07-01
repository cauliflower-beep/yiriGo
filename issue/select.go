package main

import (
	"fmt"
	"sync"
)

/*
select 语句用于监控并选择一组case语句执行相应的代码。
他看起来类似于 switch 语句，但是所有 case 中的表达式都必须是 channel 的发送或接收操作。
*/
var select_wg sync.WaitGroup

func selectDemo(ch1, ch2 chan int) {
	select {
	case <-ch1:
		fmt.Println("ch1 ...")
	case <-ch2:
		fmt.Println("ch2 ...")
	}
	select_wg.Done()
}

func main() {
	select_wg.Add(1)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go selectDemo(ch1, ch2)
	ch1 <- 1

	select_wg.Wait()

}
