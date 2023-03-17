package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	/*
		如果不关闭，当遍历到channel为空的时候，会产生死锁
		同时，调用close方法关闭通道时，会给所有等待读取通道数据的协程发送消息，这是一个非常有用的特性
			val,status := <-c
	*/
	close(c)
}

func main() {
	// 2.声明一个channel
	//c := make(chan int, 10)
	c1 := make(chan int)
	o := make(chan bool)

	// 3.Buffered Channels
	// 以下channel大小改为2及以上可以正常运行，若改为1将会报错。因为程序会阻塞在 c <- 2这一行导致死锁
	//c <- 1
	//c <- 2
	//fmt.Println(<-c, <-c)

	// 4.for-range遍历
	//go fibonacci(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	// 5.select
	//c <- 100
	//c1 <- 200 //报错！！无缓冲通道的读写必须位于不同的协程中！
	//select {
	//case i := <-c:
	//	fmt.Println(i)
	//case i := <-c1:
	//	fmt.Println(i)
	//default:
	//	fmt.Println("all channels are empty!")
	//}

	// 6.超时
	go func() {
		for {
			select {
			case v := <-c1:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout!")
				o <- true
				break
			}
		}
	}()
	<-o

}
