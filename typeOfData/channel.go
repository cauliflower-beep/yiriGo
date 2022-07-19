package main

import "sync"

/*
channel 是golang在语言级别上提供的goroutine间的通讯方式。可以用它在多个goroutine之间传递消息。

如果说goroutine是go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

golang的并发模型是 CSP ，提倡通过通信共享内存，而不是通过共享内存而实现通信，

go中的管道 channel 是一种特殊的类型，它像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序。
每一个管道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型
*/

/*
管道的声明：
1.管道是一种引用类型；
2.声明管道时需要声明类型：var i chan int；
3.声明的管道需要使用make函数分配内存后才能使用。
*/
/**********************************************************************************************/

/*
	channel支持for-range的方式进行遍历，但需要注意两个细节：
	1.遍历的时候，如果channel没有关闭，则会出现deadlock的错误;
	2.遍历的时候，如果channel已经关闭，则会正常遍历数据，遍历完后会退出遍历
*/

/**********************************************************************************************/

/*
如何优雅的关闭通道？
1.没有简单通用的方法(或内建的检查)去检查一个通道是否已经关闭，不需要改变通道状态的那种；
2.关闭一个已经关闭的通道会panic，所以如果不知道是否通道已经关闭，去关闭它是很危险的；
3.往一个已经关闭的通道发送值会panic，所以如果不知道通道是否已经关闭，往通道发送值是很危险的。

通道关闭原则
使用channel的一个基本原则就是：不要在接收方关闭channel；以及如果有多个并发发送方的话，在发送方也不能关闭channel。
换句话说，我们仅应该在一个发送goroutine中关闭通道，如果这个发送者是channel唯一的发送方的话。
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
	mc.once.Do(func() {
		close(mc.C)
	})
}
