文章地址：

https://mp.weixin.qq.com/s/d0olIiZgZNyZsO-OZDiEoA

## 内存泄漏

关于Go的内存泄露有句很有名的一句话：

`10次内存泄露，有9次是goroutine泄露。`

本文**主要介绍Go程序的goroutine泄露，掌握了如何定位和解决goroutine泄露，就掌握了内存泄露的大部分场景**。

## go pprof

pprof是Go的性能工具，可以辅助定位goroutine泄漏。

### 什么是pprof

在程序运行过程中，pprof可以记录程序的运行信息，可以是CPU使用情况、内存使用情况、goroutine运行情况等，当需要性能调优或者定位Bug时候，这些记录的信息是相当重要。

### pprof基本使用

使用pprof有多种方式，Go已经现成封装好了1个：`net/http/pprof`，使用简单的几行命令，就可以开启pprof，记录运行信息，并且提供了Web服务，能够通过浏览器和命令行2种方式获取运行数据。

可以看个简单的例子，位于demo.go中。

#### 浏览器方式

运行demo.go之后，输入网址`ip:port/debug/pprof/`打开pprof主页，从上到下依次是**5类profile信息**：

1. **block**：goroutine的阻塞信息，本例就截取自一个goroutine阻塞的demo，但block为0，没掌握block的用法
2. **goroutine**：所有goroutine的信息，下面的`full goroutine stack dump`是输出所有goroutine的调用栈，是goroutine的debug=2，后面会详细介绍。
3. **heap**：堆内存的信息
4. **mutex**：锁的信息
5. **threadcreate**：线程信息

本文主要关注goroutine和heap，这两个都会打印调用栈信息，goroutine里面还会包含goroutine的数量信息，heap则是内存分配信息，本文用不到的地方就不展示了，最后推荐几篇文章大家去看。

#### 命令行方式

连接在服务器终端上的时候，是没有浏览器可以使用的，Go提供了命令行的方式，能够获取以上5类信息，这种方式用起来更方便。

使用命令`go tool pprof url`可以获取指定的profile文件，此命令会发起http请求，然后下载数据到本地，之后进入交互式模式，就像gdb一样，可以使用命令查看运行信息，以下是5类请求的方式：

```go
# 下载cpu profile，默认从当前开始收集30s的cpu使用情况，需要等待30s
go tool pprof http://localhost:6060/debug/pprof/profile   # 30-second CPU profile
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=120     # wait 120s

# 下载heap profile
go tool pprof http://localhost:6060/debug/pprof/heap      # heap profile

# 下载goroutine profile
go tool pprof http://localhost:6060/debug/pprof/_goroutine # _goroutine profile

# 下载block profile
go tool pprof http://localhost:6060/debug/pprof/block     # _goroutine blocking profile

# 下载mutex profile
go tool pprof http://localhost:6060/debug/pprof/mutex
```

上面的`pprof/demo.go`太简单了，如果去获取内存profile，几乎获取不到什么，换一个Demo（demo2.go）进行内存profile的展示.

## 什么是内存泄漏

内存泄露指的是程序运行过程中已不再使用的内存，没有被释放掉，导致这些内存无法被使用，直到程序结束这些内存才被释放的问题。

Go虽然有GC来回收不再使用的堆内存，减轻了开发人员对内存的管理负担，但这并不意味着Go程序不再有内存泄露问题。在Go程序中，如果没有Go语言的编程思维，也不遵守良好的编程实践，就可能埋下隐患，造成内存泄露问题。

## 怎么发现内存泄漏

在Go中发现内存泄露有2种方法，一个是通用的监控工具，另一个是go pprof：

1. **`监控工具`**：固定周期对进程的内存占用情况进行采样，数据可视化后，根据内存占用走势（持续上升），很容易发现是否发生内存泄露
2. **`go pprof`**：适合没有监控工具的情况，使用Go提供的pprof工具判断是否发生内存泄露。

### 监控工具查看进程内在占用情况

**如果使用云平台部署Go程序**，云平台都提供了内存查看的工具，可以查看OS的内存占用情况和某个进程的内存占用情况，比如阿里云，我们在1个云主机上只部署了1个Go服务，所以OS的内存占用情况，基本是也反映了进程内存占用情况，OS内存占用情况如下，可以看到**随着时间的推进，内存的占用率在不断的提高，这是内存泄露的最明显现象**：

![img\\](img\1649229520(1).jpg)

注意！重点不是一开始内存占有多大，而是随着时间的推移，占用内存不断增加！

**如果没有云平台这种内存监控工具，可以制作一个简单的内存记录工具。**

1. 建立一个脚本`prog_mem.sh`，获取进程占用的物理内存情况，脚本内容如下：

   ```shell
   #!/bin/bash
   prog_name="your_programe_name"
   prog_mem=$(pidstat  -r -u -h -C $prog_name |awk 'NR==4{print $12}')
   time=$(date "+%Y-%m-%d %H:%M:%S")
   echo $time"\tmemory(Byte)\t"$prog_mem >>~/record/prog_mem.log
   ```

2. 然后使用`crontab`建立定时任务，每分钟记录1次。使用`crontab -e`编辑crontab配置，在最后增加1行：

   ```shell
   */1 * * * * ~/record/prog_mem.sh
   ```

3. 脚本输出的内容保存在`prog_mem.log`，只要大体浏览一下就可以发现内存的增长情况，判断是否存在内存泄露。如果需要可视化，可以直接黏贴`prog_mem.log`内容到Excel等表格工具，绘制内存占用图。

   ![](img\1649229784(1).jpg)

### go pprof发现存在内存问题

如果你Google或者百度，Go程序内存泄露的文章，它总会告诉你使用**pprof heap**，能够生成漂亮的调用路径图，火焰图等等，然后你根据调用路径就能定位内存泄露问题，我最初也是对此深信不疑，尝试了若干天后，只是发现内存泄露跟某种场景有关，根本找不到内存泄露的根源，**如果哪位朋友用heap就能定位内存泄露的线上问题，麻烦介绍下**。

后来读了Dave的《High Performance Go Workshop》，刷新了对heap的认识，内存pprof的简要内容如下：

![](img\1649229940(1).jpg)

文章主要讲了以下几点：

1. **内存profiling记录的是堆内存分配的情况，以及调用栈信息**，并不是进程完整的内存情况，猜测这也是在go pprof中称为heap而不是memory的原因。
2. **栈内存的分配是在调用栈结束后会被释放的内存，所以并不在内存profile中**。
3. 内存profiling是基于抽样的，默认是每1000次堆内存分配，执行1次profile记录。
4. 因为内存profiling是基于抽样和它跟踪的是已分配的内存，而不是使用中的内存，（比如有些内存已经分配，看似使用，但实际以及不使用的内存，比如内存泄露的那部分），所以**不能使用内存profiling衡量程序总体的内存使用情况**。
5. **Dave个人观点：使用内存profiling不能够发现内存泄露**。

基于目前对heap的认知，网上有两个观点：

1. **heap能帮助我们发现内存问题，但不一定能发现内存泄露问题**，这个看法与Dave是类似的。heap记录了内存分配的情况，我们能通过heap观察内存的变化，增长与减少，内存主要被哪些代码占用了，程序存在内存问题，这只能说明内存有使用不合理的地方，但并不能说明这是内存泄露。
2. **heap在帮助定位内存泄露原因上贡献的力量微乎其微**。如第一条所言，能通过heap找到占用内存多的位置，但这个位置通常不一定是内存泄露，就算是内存泄露，也只是内存泄露的结果，并不是真正导致内存泄露的根源。

下面介绍怎么用heap发现问题，然后再解释为什么heap几乎不能定位内存泄露的根因。

#### 怎么用heap发现内存问题

使用pprof的heap能够获取程序运行时的内存信息，在程序平稳运行的情况下，每个一段时间使用heap获取内存的profile，**然后使用`base`能够对比两个profile文件的差别，就像`diff`命令一样显示出增加和减少的变化**，使用一个简单的demo来说明heap和base的使用，依然使用demo2进行展示。

将demo2运行起来，执行以下命令获取profile文件，Ctrl-D退出，1分钟后再获取1次。

```go
go tool pprof http://localhost:6060/debug/pprof/heap
```

获取到两个profile文件：

```shell
$ ls
pprof.demo2.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
pprof.demo2.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
```

使用`base`把001文件作为基准，然后用002和001对比，先执行`top`看`top`的对比，然后执行`list main`列出`main`函数的内存对比，结果如下：

```go
$ go tool pprof -base pprof.demo2.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz pprof.demo2.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz

File: demo2
Type: inuse_space
Time: May 14, 2019 at 2:33pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
(pprof)
(pprof) top
Showing nodes accounting for 970.34MB, 32.30% of 3003.99MB total
      flat  flat%   sum%        cum   cum%
  970.34MB 32.30% 32.30%   970.34MB 32.30%  main.main   // 看这
         0     0% 32.30%   970.34MB 32.30%  runtime.main
(pprof)
(pprof)
(pprof) list main.main
Total: 2.93GB
ROUTINE ======================== main.main in /home/ubuntu/heap/demo2.go
  970.34MB   970.34MB (flat, cum) 32.30% of Total
         .          .     20:    }()
         .          .     21:
         .          .     22:    tick := time.Tick(time.Second / 100)
         .          .     23:    var buf []byte
         .          .     24:    for range tick {
  970.34MB   970.34MB     25:        buf = append(buf, make([]byte, 1024*1024)...) // 看这
         .          .     26:    }
         .          .     27:}
         .          .     28:
```

`top`列出了`main.main`和`runtime.main`，`main.main`就是我们编写的main函数，`runtime.main`是runtime包中的main函数，也就是所有main函数的入口，这里不多介绍了，有兴趣可以看之前的调度器文章《Go调度器系列（2）宏观看调度器》。

`top`显示`main.main` 第2次内存占用，比第1次内存占用多了970.34MB。

`list main.main`告诉了我们增长的内存都在这一行：

```go
buf = append(buf, make([]byte, 1024*1024)...)
```

001和002 profile的文件不进去看了，你本地测试下计算差值，绝对是刚才对比出的970.34MB。

#### heap“不能”定位内存泄漏

heap能显示内存的分配情况，以及哪行代码占用了多少内存，我们能轻易的找到占用内存最多的地方，如果这个地方的数值还在不断怎大，基本可以认定这里就是内存泄露的位置。

曾想按图索骥，从内存泄露的位置，根据调用栈向上查找，总能找到内存泄露的原因，这种方案看起来是不错的，但实施起来却找不到内存泄露的原因，结果是事半功倍。

原因在于一个Go程序，其中有大量的goroutine，这其中的调用关系也许有点复杂，也许内存泄露是在某个三方包里。举个栗子，比如下面这幅图，每个椭圆代表1个goroutine，其中的数字为编号，箭头代表调用关系。heap profile显示g111（最下方标红节点）这个协程的代码出现了泄露，任何一个从g101到g111的调用路径都可能造成了g111的内存泄露，有2类可能：

1. 该goroutine只调用了少数几次，但消耗了大量的内存，说明每个goroutine调用都消耗了不少内存，**内存泄露的原因基本就在该协程内部**。
2. 该goroutine的调用次数非常多，虽然每个协程调用过程中消耗的内存不多，但该调用路径上，协程数量巨大，造成消耗大量的内存，并且这些goroutine由于某种原因无法退出，占用的内存不会释放，**内存泄露的原因在到g111调用路径上某段代码实现有问题，造成创建了大量的g111**。

**第2种情况，就是goroutine泄露，这是通过heap无法发现的，所以heap在定位内存泄露这件事上，发挥的作用不大**。![](img\1649230481(1).jpg)

## goroutine泄漏怎么导致内存泄漏

### 什么是goroutine泄漏

如果你启动了1个goroutine，但并没有符合预期的退出，直到程序结束，此goroutine才退出，这种情况就是goroutine泄露。

`思考：什么会导致goroutine无法退出/阻塞？`

每个goroutine占用2KB内存，泄露1百万goroutine至少泄露`2KB * 1000000 = 2GB`内存，为什么说至少呢？

goroutine执行过程中还存在一些变量，如果这些变量指向堆内存中的内存，GC会认为这些内存仍在使用，不会对其进行回收，这些内存谁都无法使用，造成了内存泄露。

所以goroutine泄露有2种方式造成内存泄露：

1. goroutine本身的栈所占用的空间造成内存泄露。
2. goroutine中的变量所占用的堆内存导致堆内存泄露，这一部分是能通过heap profile体现出来的。

Dave在文章中也提到了，如果不知道何时停止一个goroutine，这个goroutine就是潜在的内存泄露：

```go
7.1.1 Know when to stop a goroutine

If you don’t know the answer, that’s a potential memory leak as the goroutine will pin its stack’s memory on the heap, as well as any heap allocated variables reachable from the stack.
```

### 怎么确定是goroutine泄漏引发的内存泄漏

掌握了前面的pprof命令行的基本用法，很快就可以确认是否是goroutine泄露导致内存泄露，如果你不记得了，马上回去看一下[go pprof基本知识](#go pprof基本知识)。

**判断依据：在节点正常运行的情况下，隔一段时间获取goroutine的数量，如果后面获取的那次，某些goroutine比前一次多，如果多获取几次，是持续增长的，就极有可能是goroutine泄露**。

goroutine导致内存泄露的demo3.编译运行，然后使用`go tool pprof`获取gorourine的profile文件。

```go
go tool pprof http://localhost:6060/debug/pprof/_goroutine
```

已经通过pprof命令获取了2个goroutine的profile文件:

```go
$ ls
/home/ubuntu/pprof/pprof.leak_demo.goroutine.001.pb.gz
/home/ubuntu/pprof/pprof.leak_demo.goroutine.002.pb.gz
```

同heap一样，我们可以使用`base`对比2个goroutine profile文件：

```go
$go tool pprof -base pprof.leak_demo.goroutine.001.pb.gz pprof.leak_demo.goroutine.002.pb.gz

File: leak_demo
Type: goroutine
Time: May 16, 2019 at 2:44pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
(pprof) top
Showing nodes accounting for 20312, 100% of 20312 total
      flat  flat%   sum%        cum   cum%
     20312   100%   100%      20312   100%  runtime.gopark
         0     0%   100%      20312   100%  main.alloc2
         0     0%   100%      20312   100%  main.alloc2.func1
         0     0%   100%      20312   100%  runtime.chansend
         0     0%   100%      20312   100%  runtime.chansend1
         0     0%   100%      20312   100%  runtime.goparkunlock
(pprof)
```

可以看到运行到`runtime.gopark`的goroutine数量增加了20312个。再通过002文件，看一眼执行到`gopark`的goroutine数量，即挂起的goroutine数量：

## 定位goroutine泄漏的2种方法

### web可视化查看

### 命令行交互式方法

## 总结

### goroutine泄漏的本质

### goroutine泄漏的发现和定位

### goroutine泄漏的场景

### 编码goroutine泄漏的建议

## 推荐阅读





