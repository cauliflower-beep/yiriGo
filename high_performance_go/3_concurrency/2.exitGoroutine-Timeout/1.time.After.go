package main

import (
	"fmt"
	"time"
)

/*
	超时控制在网络编程中是非常常见的，利用 context.WithTimeout 和 time.After 都能够很轻易地实现。
*/

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) error {
	/*
		这段代码的问题其实是比较明显的，具体现象见单元测试的结果。
		done是一个无缓冲channel，如果没有超时，f会向done发送信号，select会接收到done的信号，f及子协程可以正常退出
		可一旦超时，select接到time.After的信号就返回了。
		done没有了接收方receiver,而f在1s后向done发送信号。
		由于没有接收者且无缓存区，发送者sender会一直阻塞，导致协程不能退出
	*/
	done := make(chan bool)
	// 启动子协程执行函数f 结束后向 done 发送结束信号
	go f(done)
	select { // 阻塞等待 done 或者 time.After 的消息
	case <-done:
		// 没超时返回nil
		fmt.Println("done")
		return nil
	// 利用time.After 启动一个异步定时器，返回一个channel, 当超过指定的时间之后，该channel会接到信号。
	case <-time.After(time.Millisecond):
		// 超时则返回错误
		return fmt.Errorf("timeout")
	}
}

/*
	避免上述问题的出现，通常有以下几种做法
*/

// timeoutWithBuffer
//  @Description: 创建有缓冲的channel
func timeoutWithBuffer(f func(chan bool)) error {
	// 缓冲区设置为1 即使没有receiver, sender也不会阻塞
	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("goroutine done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

// doGoodthing
//  @Description: 在子协程的处理函数里面，用select发送
//  @param done
func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	/*
		使用select尝试向信道done(无缓冲)发送信号。
		如果失败了，说明缺少接收者receiver，即超时了，此时直接退出就好。
	*/
	select {
	case done <- true:
	default:
		return
	}
}

/*
	还有一些更复杂的场景，
	例如将任务拆分为多段，只检测第一段是否超时，若没有超时，后续任务继续执行，超时则终止。

	这种场景在实际的业务中更为常见.
	例如我们将服务端接收请求后的任务拆分为 2 段，一段是执行任务，一段是发送结果。那么就会有两种情况：
		1.任务正常执行，向客户端返回执行结果。
		2.任务超时执行，向客户端返回超时。
	这种情况下，就只能够使用 select，而不能能够设置缓冲区的方式了。
	因为如果给信道 phase1 设置了缓冲区，phase1 <- true 总能执行成功，那么无论是否超时，都会执行到第二阶段,而没有即时返回，这是我们不愿意看到的。
	对应到上面的业务，就可能发生一种异常情况，向客户端发送了 2 次响应：
		任务超时执行，向客户端返回超时，一段时间后，向客户端返回执行结果。
	缓冲区不能够区分是否超时了，但是 select 可以（没有接收方，信道发送信号失败，则说明超时了）。
*/

// do2phases
//  @Description: 分段任务
//  @Description: 这段代码是想表达，当父协程退出之后，子协程能不能在一定时间后也退出
func do2phases(phase1, done chan bool) {
	time.Sleep(time.Second) // 第1段
	select {
	case phase1 <- true: // 尝试向phase1信道写入信号，若成功，则说明 主go程没有超时退出，phase1存在接收方。
	default: // 因为是无缓冲通道，若主go程超时退出，没有phase1的接收方，就会执行这里，返回
		return
	}
	time.Sleep(time.Second) // 第2段
	done <- true
}

func timeoutFirstPhase() error {
	phase1 := make(chan bool)
	done := make(chan bool)
	go do2phases(phase1, done)
	select {
	case <-phase1:
		<-done
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

/*
	强制kill goroutine 可能吗？——答案是不能。
	上面的例子，即时超时返回了，但是子协程仍在继续运行，直到自己退出。那么有可能在超时的时候，就强制关闭子协程吗？
		——答案是不能。goroutine 只能自己退出，而不能被其他 goroutine 强制关闭或杀死。
	goroutine 被设计为不可以从外部无条件地结束掉，只能通过 channel 来与它通信。
	也就是说，每一个 goroutine 都需要承担自己退出的责任。
	(A goroutine cannot be programmatically killed. It can only commit a cooperative suicide.)

	详细可以参看github上关于这道题目的讨论：
	https://github.com/golang/go/issues/32610
	摘抄其中几个比较有意思的观点如下：
	1.杀死一个 goroutine 设计上会有很多挑战，当前所拥有的资源如何处理？堆栈如何处理？defer 语句需要执行么？
	2.如果允许 defer 语句执行，那么 defer 语句可能阻塞 goroutine 退出，这种情况下怎么办呢？

	因为 goroutine 不能被强制 kill，在超时或其他类似的场景下，为了 goroutine 尽可能正常退出，建议如下：

	尽量使用非阻塞 I/O（非阻塞 I/O 常用来实现高性能的网络库），阻塞 I/O 很可能导致 goroutine 在某个调用一直等待，而无法正确结束。
	业务逻辑总是考虑退出机制，避免死循环。
	任务分段执行，超时后即时退出，避免 goroutine 无用的执行过多，浪费资源。
*/

// 对于chan只需要做一次信号传递，且没有数据传输时，可以用 make(chan struct{})传递时可以用close，标准库的context就是这么玩的。
// fmt.Println(unsafe.Sizeof(struct{}{}))
