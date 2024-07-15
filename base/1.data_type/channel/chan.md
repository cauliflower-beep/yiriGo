## 概念

goroutine 运行在相同的地址空间，所以访问共享内存必须做好同步，而golang的并发模型是 CSP ，提倡通过通信共享内存，而不是通过共享内存而实现通信。

channel 是golang在语言级别上提供的goroutine间的通讯方式。可以用它在多个goroutine之间传递消息。

如果说goroutine是go程序并发的执行体，channel就是它们之间的连接，channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

本质上说， channel 是一种特殊的类型。它像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序。每一个channel都是一个**具体类型**的导管，也就是声明channel的时候需要为其指定元素类型。

channel 可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。

## 声明

管道是一种引用类型。

声明管道时需声明类型：

```go
var i chan int
var cr chan<- string //只读通道
var cw <-chan bool //只写通道
```

声明的管道必须使用make函数分配内存后才能使用：

```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```

## 读写

### 非缓冲通道

无缓冲channel的特点是：

- 当发送goroutine向channel发送数据时，如果没有接收goroutine准备好从该channel中接收数据，那么发送操作将会被阻塞，直到接收操作准备就绪；
- 当接收goroutine准备好从channel中接收数据时，如果没有发送goroutine准备好向该channel发送数据，那么接收操作将会被阻塞，直到发送操作准备就绪。

通俗来讲，非缓冲channel接收和发送数据都是阻塞的，除非另一端已经准备好(必须有配对的操作方才能执行)。这种方式可以确保通信是同步的，即发送和接收操作必须同时准备就绪，才能完成通信，不需要显式的lock(并发安全)。来看一个示例：

```go
ch <- v    // 发送v到channel
v := <-ch  // 从ch中接收数据，并赋值给v
```

所谓阻塞，也就是如果读取（v := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（如ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。

### 有缓冲通道

与无缓冲通道相对的是有缓冲通道，即Go允许指定channel的缓冲大小（可以存储多少元素）。

 ```go
ch:= make(chan bool, 4)
 ```

上面的示例创建了可以存储4个元素的bool 型元素的channel。在这个channel 中，前4个元素可以无阻塞的写入。
当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。

## 遍历

channel支持for-range的方式进行遍历，但需要注意两个细节：

1. 如果channel没有关闭，则会出现deadlock的错误;
2. 如果channel已经关闭，则会正常遍历数据，遍历完后会退出遍历

## select

上面介绍的都是只有一个channel的情况。如果存在多个channel，可通过关键字 select 监听channel上的数据流动。

select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行；当多个channel都准备好的时候，select是随机的选择一个执行的。

select里面还有default语法，类似switch的功能。default就是当监听的channel都没有准备好的时候，默认执行的(select不再阻塞等待channel)。

select经常和 for 一起使用。

## close

虽然通道可以关闭，但并不是一个必须执行的方法。因为通道本身会通过垃圾回收器，根据它是否可以访问来决定是否回收。

> 如何优雅的关闭通道？

首先要明确一点，官方并没有给出一个简单通用的方法(或内建的检查)，在不需要改变通道状态的前提下，去判断一个通道是否已经关闭。

但可以通过语法 v, ok := <- ch测试 channel是否关闭，如果ok返回false,说明channel已经没有任何数据并且已被关闭。

不过即使存在这样的方法，它的用处也会很有限，就像用来检查通道中当前值的个数的内建函数len。因为被检查的通道的状态可能会在函数返回后瞬间改变，所以返回的那个值并没法反映那个通道的最新状态。

例如 v, ok := <- ch  ok返回true，不代表你可以安全的往那个通道发送值或者关闭它。

另外，关闭一个已经关闭的通道会panic，所以如果不知道是否通道已经关闭，去关闭它是很危险的；并且往一个已经关闭的通道发送值也会panic，所以如果不知道通道是否已经关闭，往通道发送值是很危险的。

> 通道关闭原则

概括起来主要有如下两点：

1. 不要在接收方关闭channel；以及如果有多个并发发送方的话，在发送方也不能关闭channel。换句话说，我们仅应该在一个发送goroutine中关闭通道，如果这个发送者是channel唯一的发送方的话。
2. channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的时候可以关闭。

> 简单粗暴的方案

如果无论如何都要从通道的接收方，或者多个发送方关闭一个通道的话，可以使用recover机制防止可能的panic把程序搞崩。但这个方案不仅打破了通道关闭原则，也可能在进程内发生data race。

> 礼貌的方案

可以使用`sync.Once`来关闭通道。

小节代码：[close](.\close.go).

## 超时控制

可以参考

...\high_performance_go\3_concurrency\2.exitGoroutine-Timeout

中关于goroutine超时控制的讲解。

## deadlock

