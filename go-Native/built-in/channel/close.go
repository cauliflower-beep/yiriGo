package main

import "sync"

/***********************************************************************************************/
/*
虽然通道可以关闭，但并不是一个必须执行的方法。因为通道本身会通过垃圾回收器，根据它是否可以访问来决定是否回收。

如何优雅的关闭通道？
1.没有简单通用的方法(或内建的检查)去检查一个通道是否已经关闭，不需要改变通道状态的那种；
	可以通过语法 v, ok := <- ch测试 channel是否关闭，如果ok返回false,说明channel已经没有任何数据并且已被关闭
	即使有一个简单的内建closed函数能检查一个通道是否已关闭，它的用处也会很有限，就像用来检查通道中当前值的个数的内建函数len。
这是因为被检查的通道的状态可能会在函数返回后瞬间改变，所以返回的那个值并没法反映那个通道的最新状态.
	例如 v, ok := <- ch  ok返回true，不代表你可以安全的往那个通道发送值或者关闭它。
2.关闭一个已经关闭的通道会panic，所以如果不知道是否通道已经关闭，去关闭它是很危险的；
3.往一个已经关闭的通道发送值会panic，所以如果不知道通道是否已经关闭，往通道发送值是很危险的。

通道关闭原则：
1.不要在接收方关闭channel；以及如果有多个并发发送方的话，在发送方也不能关闭channel。
换句话说，我们仅应该在一个发送goroutine中关闭通道，如果这个发送者是channel唯一的发送方的话。
2.channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的时候可以关闭

简单粗暴的方案：
如果你无论如何都要从通道的接收方,或者多个发送方关闭一个通道的话，可以使用recover机制防止可能的panic把程序搞崩。
这个方案不仅打破了通道关闭原则，也可能在进程内发生data race

礼貌的方案：
很多人喜欢用 sync.Once 来关闭通道
*/

type T int

// safeClose 简单粗暴的方案关闭可能关闭的通道
func safeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// the return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch) // panic if ch is closed
	return true
}

// safeSend 简单粗暴的方案向一个可能关闭的channel中发送数据
func safeSend(ch chan T, val T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- val // panic if ch is closed
	return false
}

// Mychannel 礼貌的关闭channel——sync.Once 版
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
