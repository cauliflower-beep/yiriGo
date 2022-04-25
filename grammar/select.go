package main

import (
	"fmt"
	"time"
)

/*
select用于监控并选择一组 case 语句执行相应的代码；
看起来类似于 switch 语句，但是select语句中所有case中的表达式都必须是channel的发送或接收操作；
一个典型的使用示例如下：
 */

func waitCh (ch1,ch2 chan string){
	select{
	case <- ch1:
		fmt.Println("ch1: lufy")
	case ch2 <- "goku":
		fmt.Println("ch2: goku")
	}
}

func main(){
	tricker := time.NewTicker(time.Second)
	ch1 := new(chan string)
	ch2 := new(chan string)
	for{
		<- tricker.C
		now := time.Now()
		if now.Second() %2 ==0{

		}
	}
}
