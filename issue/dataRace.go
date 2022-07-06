package main

import (
	"fmt"
	"sync"
)

/*
伴随着并发的使用，可能会发生可怕的数据争用(data race)问题
——只要有两个以上的goroutine并发访问同一变量，且至少其中的一个是写操作的时候，就会发生数据竞争；全是读的情况下是不存在数据竞争的。
一旦遇到 data race,由于不知道什么时候发生，将会是难以发现和调试的错误之一。
*/

/*
getNumber 函数中，先声明一个变量i;之后在goroutine中单独对i进行设置，而这时候程序也正在从函数中
返回i，由于不知道goroutine是否已完成对i值得修改，因此将会有两种情况发生：
1.goroutine先完成对i值得修改，最后返回的i值被设置为5；
2.变量 i 的值从函数返回，结果为默认的0.
根据这两种情况哪一个先完成，结果不同，这就是一个很典型的数据竞态.
*/
func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	return i
}

/*
检查竞态
Go(从V1.1开始)具有内置的数据竞争检测器，可以使用它来查明潜在的数据竞争条件。使用它就像在普通的Go命令
行中添加标志一样简单：

运行时检查竞态的命令: go run -race main.go
构建时检查竞态的命令： go build -race main.go
测试时检查竞态的命令： go test -race main.go

所有避免产生竞态背后的核心原则都是防止对同一变量或内存位置同时进行读写和访问。
*/

/*
避免竞态的方式
go为避免竞态提供了很多解决方案。所有这些方案都有助于确保如果我们正在写入变量，则该变量的访问将被阻止。
1)WaitGroup等待：解决数据竞态的最直接方法是阻止读取访问，直到写操作完成为止，这时候可以用WaitGroup来实现；
  注意：这种方式使用的时候必须保证Add和Done方法出现的次数一致，最后调用Wait等待添加的任务都执行完毕。如果
	   Add和Done数量不一致，就会一直阻塞程序，无限制消耗内存等资源，知道资源耗尽，服务宕机.
2)channel阻塞等待：该方法在原则上与上一种方法类似，只是使用的是channel而不是等待组；
3)返回channel通道：除了使channel阻塞等待，还可以返回一个channel，一旦获得结果，就可以通过该通道直接推送结
  果。此方法本身不会进行任何阻塞，相反，它保留了阻塞调用代码的时机；这种方式也更加灵活，因为它允许更高级别的功
  能决定自己的阻塞和并发机制，而不是将getNumber功能视为同步功能；
4)使用互斥锁：上面3种方式解决的是i在写操作完成之后才能读取的情况。现在有以下条件：
	不管读写顺序如何，只要求它们不能同时发生
  针对这种场景，应该考虑使用互斥锁。
*/

// WaitGroup
func getNumberByWg() int {
	var i int
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		i = 5
		wg.Done()
	}()
	wg.Wait()
	return i
}

// Channel1
func getNumberChan() int {
	var i int
	done := make(chan struct{})
	go func() {
		i = 5
		done <- struct{}{}
	}()
	<-done // 阻塞等待
	return i
}

// Channel2
func getNumberByChan() int {
	num := make(chan int)
	go func() {
		num <- 5
	}()
	return <-num
}

// mutex Todo

func main() {
	//fmt.Println(getNumber())
	fmt.Println(getNumberByWg())
	fmt.Println(getNumberChan())
	fmt.Println(getNumberByChan())

}
