## 互斥锁

互斥锁是一种常用的控制共享资源访问的方法，Go中使用`sync`包的`Mutex`类型来实现互斥锁。它能够保证同一时间有且只有一个`goroutine`可以进入临界区，访问共享资源，其他的goroutine则在等待锁。

多个goroutine同时等待一个锁时，唤醒的策略是随机的。

互斥锁的问题在于，在读的时候，也会阻塞下一个读线程，读多写少的情况下效率很低。

### 方法

`mutex`提供了几个方法：

```go
func (*Mutex) Lock //锁定

func (*Mutex) Unlock // 解锁
/*
在Go1.18新提供了TryLock()方法可以非阻塞式的取锁操作：
tryLock()：调用TryLock方法尝试获取锁;
当锁被其他 goroutine 占有，或者当前锁正处于饥饿模式，它将立即返回 false;
当锁可用时尝试获取锁，获取失败不会自旋/阻塞，也会立即返回false；
 */
```

### 实现
mutex的结构比较简单，只有两个字段：
```go
type Mutex struct{
	state int32 // 表示当前互斥锁的状态，复合型字段；
	seme uint32 // 信号量变量，用来控制等待goroutine的阻塞、休眠和唤醒
}
```
互斥锁本身是个复杂的东西，之所以两个字段就可以实现，得益于它优秀的设计方式——使用位做标志。

state不同位表示了不同的状态，使用最小的内存来表示更多的意义。其中低三位由低到高分别表示mutexLocked、mutexWoken 和 mutexStarving.剩下的位则用来表示当前有多少个goroutine在等待锁：

```go
const (
   mutexLocked = 1 << iota // 表示互斥锁的锁定状态
   mutexWoken // 表示从正常模式被从唤醒
   mutexStarving // 当前的互斥锁进入饥饿状态
   mutexWaiterShift = iota // 当前互斥锁上等待者的数量
)
```
mutex最开始的实现只有正常模式，该模式下等待的线程按照先进先出的方式获取锁。但是新创建的goroutine会与刚被唤起的goroutine竞争，可能导致刚被唤起的goroutine获取不到锁，长时间被阻塞下去，所以Go在1.9中进行了优化，
引入了饥饿模式。

当goroutine超过1ms没有获取到锁，就会将当前互斥锁切换到饥饿模式，在饥饿模式中，互斥锁会直接交给等待队伍最前面的goroutine，新的goroutine在该状态下不能获取锁，也不会进入自旋状态，他们只会在队列的末尾等待。如果一个goroutine
获得了互斥锁并且他在队列的末尾或者它等待的时间少于1ms，当前的互斥锁就会切换为正常模式。

https://mp.weixin.qq.com/s/rRPQ6YN15P7UODe1b7oyAA
## 读写锁

很多实际场景下是读多写少的，当我们并发的去读取一个资源，不涉及资源修改的时候，是没必要加锁的。这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用`sync`包中的`RWMutex`实现。

读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的`goroutine`如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个`goroutine`获取写锁之后，其他的`goroutine`无论是获取读锁还是写锁都会等待。

`RWMutex`提供了四个方法：

```go
func (*RWMutex) Lock // 写锁定

func (*RWMutex) Unlock // 写解锁

func (*RWMutex) RLock // 读锁定

func (*RWMutex) RUnlock // 读解锁
```

需要注意的是读写锁适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。

### 读锁可以不加吗？

既然并发场景中读锁可以共享，那跟不加锁有什么区别？可以不加吗？

当然不可以。读锁是读的时候共享，在读-写的时候是互斥的。换言之，共享的读锁是为了锁住写线程，在读的时候不能写。