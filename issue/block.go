package main

import (
	"fmt"
	"time"
)

/*
某些情况下，我们希望程序阻塞在某一行，来观察一些数据变化或者运行效果
例如在主go程中，我们希望主go程可以阻塞等待其他go程执行结束
下面介绍几种方法来实现阻塞的效果
 */

/*
1.sync.WaitGroup
 */

/*
2.空 select{}
select{}是一个没有任何 case 的 select，它会一直阻塞
 */

/*
3.死循环 for{}
虽然能阻塞，但会100%占用cpu,不建议使用
 */

/*
4.sync.Mutex

 */
func talk(){
	for i:=0;i<10;i++{
		fmt.Println("how are you?")
		time.Sleep(time.Second)
	}
}

func main(){
	go talk()
	select{}
}
