## 概念

`golang`中的协程可以理解为**用户级线程**，它是对内核透明的，即系统并不知道有协程的存在，是完全由用户自己的程序调度的，依赖于go语言运行时自身提供的调度器。

`golang`一大特色就是从语言层面原生支持协程，在函数或者方法前加 `go` 关键字就可以创建一个协程。可以说golang中的协程就是goroutine.

go运行时的调度器使用 `runtime.GOMAXPROCS()` 函数来确定需要使用多少个 OS 线程来**同时**执行go代码。

Go1.5版本之前，默认使用的是单核心执行，之后默认值是机器上的**CPU核心数**。例如在一个8核心的机器上，调度器会把go代码同时调度到8个os线程上跑。

本小节代码：[base](.\base.go)

## 特性

goroutine的大小并不是一成不变的，并且每个goroutine需要能够独立运行，所以他们要有自己独立的栈。

假如每个goroutine分配固定的栈大小且不能增长，太小则会导致溢出，太大又会浪费空间，无法存在许多goroutine。

为了解决这个问题，goroutine初始时只给栈分配很小的空间，该空间大小随着使用过程自动增长。这也是为什么Go可以开千千万万个goroutine而不会耗尽内存。

具体怎么实现的呢?

每次执行函数调用时，go的runtime都会进行检测。若当前栈的大小不够用，则会触发“终中断”，从当前函数进入到go的运行时库。go的运行时库会保存此时的函数上下文环境，然后分配一个新的足够大的栈空间，将旧栈的内容拷贝到新栈中，并做一些设置，使得当函数恢复运行时，函数会在新分配的栈中继续执行，
仿佛整个过程都没发生过一样。函数视角来看，会觉得自己使用的是一块儿“无限大小”的栈空间。

计算一个goroutine的大小：[size](.\size.go).

## 生命周期

main函数结束的时候，所有goroutine都会跟着一起结束。这个是老生常谈。否则我们也不会在各种goroutine初始教程里，看到`time.sleep`这种让 main 等待goroutine执行完成的操作。

但main函数未结束时，父goroutine以及由它发起的子goroutine的生命周期就是另外一套逻辑了：

- **父函数**结束，子goroutine仍会继续执行；
- **父goroutine**结束，子goroutine仍会继续执行；

简而言之，除`main主go程`外，其他情况下的父子协程相互独立。项目中要格外注意父子协程的退出关系，以免造成大量的孤儿协程。

本小节代码：[lifeCycle](.\lifeCycle.go)

## 优雅关闭子go程

### context

参考 [_context](..\..\4.pkg\_context)

### channel

todo.

## 限定go程数量

> goroutine可以无限开启吗？

goroutine体积轻量，并且得益于优质的**GMP调度**，使得我们可以轻易创建并高效运行一组goroutine。

那goroutine是否可以无限开辟呢？如果做一个服务器或者处理一些高并发业务的场景，能否随意的开辟 goroutine并且放养不管呢？让他们自生自灭，毕竟有强大的GC和优质的调度算法支撑？

答案是不行的，参考本小节代码：[cntUnlimit](.\cntUnlimit.go).

运行到一定时间，终端会输出：
```go
panic: too many concurrent operations on a single file or socket (max 1048575)
```

并被操作系统以`kill`信号强制终结该进程。

所以，我们迅速开辟的goroutine（不控制并发数量）会在短时间内占据操作系统的资源（cpu、内存、文件描述符等）。整个过程发展如下：

- CPU使用率浮动上涨；
- Memory占用不断上涨；
- 主进程崩溃(被kill掉)。

这些资源实际上是所有用户态程序**共享**的资源，所以大批的goroutine最终引发的灾难不仅仅是自身， 还会关联其他运行的程序。 所以在编写逻辑业务的时候，限制goroutine是我们必须要重视的问题。

> 如何限定goroutine数量?

这里列举一些简单的方法。

*方法一：只使用有buffer的channel来限制*

小节代码：[cntLimitWithBufferChan](.\cntLimitWithBufferChan.go)，运行结果：
```go
...
go func 352277 goroutine count = 4
go func 352278 goroutine count = 4
go func 352279 goroutine count = 4
go func 352280 goroutine count = 4
...
```

从结果看，程序并没有崩溃，而是按部就班的顺序执行，并且go的数量控制在了3，（4的原因是还有个main goroutine）。

![image-20240715161159479](.\imgs\chan限制go程数量.png)

使用buffer为3的channel，实际上是限制了for循环的速度，也就是创建go程的速度。而go程的结束速度取决于`handler`函数的执行速度。这样我们就能保证同一时间内运行的goroutine数量与buffer的数量一致，达到限定效果。

*方法二：channel与sync同步组合*

小节代码：[cntLimitWithSync&chan](.\cntLimitWithSync&Chan.go)

比起方法一，这种方式可以避免因主线程结束而导致的子goroutine未全量创建的问题。

*方法三：使用无缓冲channel与任务发送/执行分离方式*

小节代码：[cntLimitWithChanNoBuffer](.\cntLimitWithChanNoBuffer.go)

这里实际上是将任务的发送和执行做了业务上的分离，并将输入的task平均分散给固定数量的goroutine来执行。输入的SendTask频率可控，执行Goroutine的数量也可控。这也是很多Go框架的worker工作池的最初设计理念。

![image-20240715163957737](.\imgs\固定数量的goroutine池.png)



