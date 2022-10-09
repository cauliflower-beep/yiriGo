package main

import (
	"fmt"
	"sync"
	"time"
)

/*1.什么是channel?
goroutine 运行在相同的地址空间，因此访问共享内存必须做好同步。
golang的并发模型是 CSP ，提倡通过通信共享内存，而不是通过共享内存而实现通信。
channel 是golang在语言级别上提供的goroutine间的通讯方式。可以用它在多个goroutine之间传递消息。

如果说goroutine是go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

go中的管道 channel 是一种特殊的类型，它像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序。
每一个管道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型；
channel 可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。
*/

/*
管道的声明：
1.管道是一种引用类型；
2.声明管道时需要声明类型：
	var i chan int
	var cr chan<- string //只读通道
	var cw <-chan bool //只写通道
3.声明的管道必须使用make函数分配内存后才能使用：
	ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})
*/
/**********************************************************************************************/
/*2.channel 的读写：
	ch <- v    // 发送v到channel
	v := <-ch  // 从ch中接收数据，并赋值给v
非缓冲channel接收和发送数据都是阻塞的，除非另一端已经准备好(必须有配对的操作方才能执行)。
这样就使得Goroutines同步变的更加的简单，而不需要显式的lock(并发安全)。
所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
无缓冲channel是在多个goroutine之间同步很棒的工具。
*/
/***********************************************************************************************/
/*3.有缓冲通道Buffered Channels
上面介绍的是默认的非缓冲类型的channel，不过Go也允许指定channel的缓冲大小，就是channel可以存储多少元素。
	ch:= make(chan bool, 4)
创建了可以存储4个元素的bool 型channel。在这个channel 中，前4个元素可以无阻塞的写入。
当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
*/
/***********************************************************************************************/
/*4.channel遍历
channel支持for-range的方式进行遍历，但需要注意两个细节：
	1.遍历的时候，如果channel没有关闭，则会出现deadlock的错误;
	2.遍历的时候，如果channel已经关闭，则会正常遍历数据，遍历完后会退出遍历
需要注意的是，虽然通道可以关闭，但并不是一个必须执行的方法。因为通道本身会通过垃圾回收器，根据它是否可以访问来决定是否回收。

如何优雅的关闭通道？
1.没有简单通用的方法(或内建的检查)去检查一个通道是否已经关闭，不需要改变通道状态的那种；
	可以通过语法 v, ok := <- ch测试 channel是否关闭，如果ok返回false,说明channel已经没有任何数据并且已被关闭
2.关闭一个已经关闭的通道会panic，所以如果不知道是否通道已经关闭，去关闭它是很危险的；
3.往一个已经关闭的通道发送值会panic，所以如果不知道通道是否已经关闭，往通道发送值是很危险的。

通道关闭原则：
1.不要在接收方关闭channel；以及如果有多个并发发送方的话，在发送方也不能关闭channel。
换句话说，我们仅应该在一个发送goroutine中关闭通道，如果这个发送者是channel唯一的发送方的话。
2.channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的时候可以关闭
*/

// 礼貌的关闭channel
type Mychannel struct {
	C    chan int
	once sync.Once
}

func NewMychannel() *Mychannel {
	return &Mychannel{C: make(chan int)}
}

func (mc *Mychannel) SafeClose() {
	// 礼貌的关闭channel
	mc.once.Do(func() {
		close(mc.C)
	})
}

/**************************************************************************************************/
/*5.select
上面介绍的都是只有一个channel的情况，如果存在多个channel，可通过关键字 select 监听channel上的数据流动。
select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行；
当多个channel都准备好的时候，select是随机的选择一个执行的。
select里面还有default语法，类似switch的功能。default就是当监听的channel都没有准备好的时候，默认执行的(select不再阻塞等待channel)。

select经常和 for 一起使用
*/
/****************************************************************************************************/
/*6.超时
有时候会出现goroutine 阻塞的情况，如何避免呢?
方案之一就是利用select设置超时。
*/
/****************************************************************************************************/
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
	// 2.声明一个管道
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
