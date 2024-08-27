package main

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

// func main() {
// 	// 2.声明一个channel
// 	// c := make(chan int, 10)
// 	c1 := make(chan int)
// 	o := make(chan bool)
//

//
// 	// 4.for-range遍历
// 	// go fibonacci(cap(c), c)
// 	// for i := range c {
// 	//	fmt.Println(i)
// 	// }
//

//
// 	// 6.超时
// 	go func() {
// 		for {
// 			select {
// 			case v := <-c1:
// 				fmt.Println(v)
// 			case <-time.After(5 * time.Second):
// 				fmt.Println("timeout!")
// 				o <- true
// 				break
// 			}
// 		}
// 	}()
// 	<-o
//
// }
