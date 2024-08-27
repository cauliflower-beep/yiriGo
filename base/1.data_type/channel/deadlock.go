package main

import "fmt"

/*
情况1：单线程的情况下，往channel中放入数据量超过其channel大小的时候
*/
func demo1() {
	ch := make(chan int, 3)
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

/*
情况2：单线程的情况下，把channel中的数据全部取出后，还继续取数据的时候
*/
func demo2() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	for {
		fmt.Println(<-ch)
	}
}

/*
情况3：channel没有关闭的情况下，使用for range进行遍历
*/
func demo3() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	// close(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
func main() {
	// demo1()
	demo2()
	// demo3()
}
