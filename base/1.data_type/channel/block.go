package main

import (
	"fmt"
	"time"
)

/************************no buffer chan******************************/
// 阻塞的情况
func blockChanNoBuffer() {

	c := make(chan int)

	c <- 5                            // 无缓冲通道，对chan的读写必须位于不同协程，否则会阻塞
	fmt.Println("yep. i am printed.") // 打印不出来这个
	select {
	case i := <-c:
		fmt.Println(i)
	default:
		fmt.Println("channel c is empty!")
	}
}

func unblockChanNoBuffer() {
	c := make(chan int)
	var i int

	go func() {
		time.Sleep(3 * time.Second)
		c <- 5
	}()
	fmt.Println("yep. i am printed.") // 能正常打印
	for {
		select {
		case i = <-c:
			fmt.Println(i)
			goto overCircle
		default:
			fmt.Println("channel c is empty!")
		}
	}
overCircle:
	fmt.Println("channel c's content is:", i)
}

/**************************buffer chan************************************/
func blockBufferChan1() {
	c := make(chan int, 1)

	c <- 5
	c <- 6                            // 单线程 通道缓冲区已满，再写就会阻塞（死锁）
	fmt.Println("yep. i am printed.") // 打印不出来这个
}

func blockBufferChan2() {
	c := make(chan int, 1)

	<-c                               // 单线程 通道缓冲区已空，继续读就会阻塞（死锁）
	fmt.Println("yep. i am printed.") // 打印不出来这个
}

func unblockBufferChan1() {
	c := make(chan int, 1)

	c <- 5                            // 单线程 缓冲区未满即可写入 不会阻塞
	fmt.Println("yep. i am printed.") // 正常打印
	select {
	case i := <-c:
		fmt.Println(i)
	default:
		fmt.Println("channel c is empty!")
	}
}

func unblockBufferChan2() {
	c := make(chan int, 1)

	c <- 5
	<-c                               // 单线程 缓冲区未空即可读出 不会阻塞
	fmt.Println("yep. i am printed.") // 正常打印
	select {
	case i := <-c:
		fmt.Println(i)
	default:
		fmt.Println("channel c is empty!")
	}
}
