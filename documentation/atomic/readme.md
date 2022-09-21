## 前言

多个并发线程操作同一块内存时，往往需要考虑并发安全的问题。

例如go中的map类型，不允许多个 goroutine 同时对其读写。

以往解决这个问题，比较直观的是加读写锁。但是加锁涉及到内核态的上下文转换，比较耗时，代价高。

如果是针对基本数据类型，我们还可以使用原子操作来保证并发的安全。

因为原子操作是通过CPU指令，也就是硬件层次去实现的，我们在用户态就可以完成，也不需要像mutex那样记录很多状态，性能比加锁操作更好。

当然mutex不只是对变量的并发控制，更多的是对代码块的并发控制，二者侧重点不一样。

## 原子操作

原子操作是指进行过程中不能被中断的操作。

针对某个共享变量的原子操作，在被执行的过程中，CPU绝不会再去进行其他的针对该值的操作。

为了实现这样的严谨性，原子操作仅会由一个独立的CPU指令代表和完成。原子操作是无锁的，常常直接通过CPU指令直接实现。

事实上，其他同步技术的实现常常依赖于原子操作。

具体的原子操作在不同的操作系统中实现是不同的。比如在Intel的CPU架构机器上，主要是使用总线锁的方式实现的。 大致的意思就是当一个CPU需要操作一个内存块的时候，向总线发送一个LOCK信号，所有CPU收到这个信号后就不对这个内存块进行操作了。 等待操作的CPU执行完操作后，发送UNLOCK信号，才结束。 在AMD的CPU架构机器上就是使用MESI一致性协议的方式来保证原子操作。 所以我们在看atomic源码的时候，我们看到它针对不同的操作系统有不同汇编语言文件。

## Go中原子操作的支持

Go语言的`sync/atomic`提供了对原子操作的支持，用于同步访问整数和指针。

- Go语言提供的原子操作都是非入侵式的
- 原子操作支持的类型包括`int32`、`int64`、`uint32`、`uint64`、`uintptr`、`unsafe.Pointer`

竞争条件是由于异步的访问共享资源，并试图同时读写该资源而导致的，使用互斥锁和通道的思路都是在线程获得到访问权后阻塞其他线程对共享内存的访问，而使用原子操作解决数据竞争问题则是利用了其不可被打断的特性。

## —ADD

## —Load

Load 方法是为了防止在读取过程中，有其他协程发起修改动作，影响了读取结果，常用于**配置项**的整个读取。

## —Store

有原子读取，就有原子修改值，前面提到过的 Add 只适用于 int、uint 类型的增减，并没有其他类型的修改，而 Sotre 方法通过 unsafe.Pointer 指针原子修改，来达到了对其他类型的修改。

## —**CompareAndSwap**

### 什么是CAS？

使用锁时，线程获取锁是一种**`悲观锁策略`**，即假设每一次执行临界区代码都会产生冲突，所以当前线程获取到锁的时候同时也会阻塞其他线程获取该锁。

而CAS操作（又称为`无锁操作`）是一种**`乐观锁策略`**，它假设所有线程访问共享资源的时候不会出现冲突，既然不会出现冲突自然而然就不会阻塞其他线程的操作。因此，线程就不会出现阻塞停顿的状态。

那么，如果出现冲突了怎么办？无锁操作是使用**`CAS(compare and swap)`**又叫做比较交换来鉴别线程是否出现冲突，出现冲突就重试当前操作，直到没有冲突为止。

### CAS的操作过程

CAS比较交换的过程可以通俗的理解为`CAS(V,O,N)`，包含三个值分别为：**`V 内存地址存放的实际值`；`O 预期的值（旧值）`；`N 更新的新值`**。当V和O相同时，也就是说旧值和内存中实际的值相同表明该值没有被其他线程更改过，即该旧值O就是目前来说最新的值了，自然而然可以将新值N赋值给V。反之，V和O不相同，表明该值已经被其他线程改过了则该旧值O不是最新版本的值了，所以不能将新值N赋给V，返回V即可。当多个线程使用CAS操作一个变量是，只有一个线程会成功，并成功更新，其余会失败。失败的线程会重新尝试，当然也可以选择挂起线程

CAS的实现需要硬件指令集的支撑，在JDK1.5后虚拟机才可以使用处理器提供的**CMPXCHG**指令实现。

### CAS的问题

- `ABA`

  因为CAS会检查旧值有没有变化，这里存在这样一个有意思的问题。比如一个旧值A变为了成B，然后再变成A，刚好在做CAS时检查发现旧值并没有变化依然为A，但是实际上的确发生了变化。解决方案可以沿袭数据库中常用的乐观锁方式，添加一个版本号可以解决。原来的变化路径A->B->A就变成了1A->2B->3C。

- `自旋时间过长`

  使用CAS时非阻塞同步，也就是说不会将线程挂起，会自旋（无非就是一个死循环）进行下一次尝试，如果这里自旋时间过长对性能是很大的消耗。如果JVM能支持处理器提供的pause指令，那么在效率上会有一定的提升。

-  `只能保证一个共享变量的原子操作`

   当对一个共享变量执行操作时CAS能保证其原子性，如果对多个共享变量进行操作,CAS就不能保证其原子性。有一个解决方案是利用对象整合多个共享变量，即一个类中的成员变量就是这几个共享变量。然后将这个对象做CAS操作就可以保证其原子性。atomic中提供了AtomicReference来保证引用对象之间的原子性。

### go中的CAS

go中的CAS操作借用了CPU提供的原子性指令来实现。

CAS操作修改共享变量时候不需要对共享变量加锁，而是通过类似乐观锁的方式进行检查，本质还是不断的占用CPU 资源换取加锁带来的开销（比如上下文切换开销）。

原子操作中的CAS(Compare And Swap),在`sync/atomic`包中，这类原子操作由名称以`CompareAndSwap`为前缀的若干个函数提供。

```go
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)

func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)

func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)

func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)

func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)

func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
```

`CompareAndSwap`函数会先判断参数addr指向的操作值与参数old的值是否相等，仅当此判断得到的结果是true之后，才会用参数new代表的新值替换掉原先的旧值，否则操作就会被忽略。

## —Swap

Swap 方法实现了对值的原子交换，不仅 int，uint 可以交换，指针也可以。

## 总结

1. atomic 很多时候可能都没有使用上，毕竟 mutex 的拓展性比较好，使用起来也比较友好。但这并不妨碍我们对极致性能的追求，有时候，细节决定了性能！