package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Pool是可伸缩、并发安全的临时对象池，用来存放已经分配但暂时不用的临时对象，通过对象重用机制，缓解GC压力，提高程序性能。

通俗来讲，它本身是一个池子，你要和他建立一个约定，之后你需要什么，就可以从中获取，用完了还可以还回去。但是下次拿的时候，
就不一定是你上次存的那个了。所以这个池子设计的目的就很明确了，就是为了复用已经使用过的对象，从而优化内存使用和回收。一
开始这个池子会初始化一些对象供你使用，如果不够了，就会通过new产生一些，当你放回去了之后，这些对象会被别人复用，当对象
特别大并且使用非常频繁的时候可以大大减少对象的创建和回收的时间。

它适合存储一些会在goroutine之间共享的临时对象，其中保存的任何项都可能随时不做通知的释放掉（当这个对象只有sync.Pool
持有时，对象内存会被释放掉），所以不适合用于存放诸如 socket 长连接或数据库连接的对象。

临时对象池几乎是对可读性影响最小且优化效果显著的手段。基本上，业内以高性能著称的开源库，都会使用到。
*/

var strPool = sync.Pool{
	/*
		Pool 本身是个结构体，里面的New字段是一个函数类型;
		可以通过new去定义这个池子里面存放的是什么东西。
		一个池子中只能存放一种类型的东西，比如这个池子中只能存放字符串。
	*/
	New: func() interface{} {
		return "Pool test"
	},
}

func main() {
	/*
		随时可以通过Get方法从池子中获取我们放入进去的数据;
		如果获取不到，就会通过New新建一个
	*/
	str := strPool.Get()
	fmt.Println(str)
	/*
		用完了之后可以通过Put方法放回去，或者放别的同类型的数据进去
	*/
	strPool.Put(123)
	strPool.Put(456)
	runtime.GC()
	runtime.GC()
	str2 := strPool.Get()
	fmt.Println(str, str2)
	str3 := strPool.Get()
	fmt.Println(str, str2, str3)
}

/*
参考:
https://www.jianshu.com/p/8fbbf6c012b2
*/
