## 背景

context翻译成中文是”上下文”，即它可以控制一组呈树状结构的goroutine，每个goroutine拥有相同的上下文。

golang在1.6.2的时候还没有自己的context，在1.7的版本中就把golang.org/x/net/context包加入到了官方的库中。golang 的 Context包，是专门用来简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

也就是说，context本质上是一种并发控制技术，应用开发中经常使用。

它与waitgroup最大的不同点是context对于派生goroutine有更强的控制力，可以控制多级goroutine.

## 为什么需要context

比如有一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine。这样的话， 我们就可以通过Context，来跟踪这些goroutine，并且通过Context来控制他们的目的。

另外一个实际例子是，在Go服务器程序中，每个请求都会有一个goroutine去处理。然而，处理程序往往还需要创建额外的goroutine去访问后端资源，比如数据库、RPC服务等。由于这些goroutine都是在处理同一个请求，所以它们往往需要访问一些共享的资源，比如用户身份信息、认证token、请求截止时间等。而如果请求超时或者被取消后，所有的goroutine都应该马上退出并且释放相关的资源。这种情况下使用waitgroup就不太方便，因为子goroutine的数量不容易确定。此时需要用Context来为我们取消掉所有goroutine

## 定义

ontext的主要数据结构是一种嵌套的结构或者说是单向的继承关系的结构，比如最初的context是一个小盒子，里面装了一些数据，之后从这个context继承下来的children就像在原本的context中又套上了一个盒子，然后里面装着一些自己的数据。或者说context是一种分层的结构，根据使用场景的不同，每一层context都具备有一些不同的特性，这种层级式的组织也使得context易于扩展，职责清晰。

context 包的核心是 struct Context，声明如下：

```go
type Context interface {

Deadline() (deadline time.Time, ok bool)

Done() <-chan struct{}

Err() error

Value(key interface{}) interface{}

}
```

