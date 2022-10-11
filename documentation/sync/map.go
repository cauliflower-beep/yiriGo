package main

import (
	"fmt"
	"sync"
)

/*
	map的问题:
		并发场景下，map有一个致命的坑点——多协程并发写可能会出现 fatal error:concurrent map read and map write 的错误；
		当然根据golang官方文档，只要有更新操作的时候(比如一个goroutine在写，多个goroutine读)，map就是非线程安全的。
		如果使用场景只是并发读，那么map并发安全
		源码分析：https://blog.csdn.net/lyn_00/article/details/108069481

	一般情况下，解决并发读写map的思路是加锁，或者把一个map切分成若干个小map，对key近行哈希。
	业界中使用最多的模式是：
		1.原生 map + 互斥锁 或者 读写锁；
		2.标准库 sync.Map(Go 1.9及之后)

	sync.Map:
		1.线性安全，读取、插入、删除都保持在常数级的时间复杂度。
		2.零值有效，并且零值是一个空map。在第一次使用之后，不允许被拷贝。
		3.大多数代码应该直接使用原生的map，而不是单独的锁或协程控制，以获得更好的类型安全性和可维护性
		4.针对以下场景做了性能优化：
			*当一个key只被写入一次但被多次读取时，例如在只会增长的缓存中；
			*当多个goroutines 读取、写入和覆盖不相干的key时
			这两种情况与Go map搭配单独的Mutex或RWMutex相比，可以大大减少锁的争夺。
*/

func main() {
	var sm sync.Map

	// 1.写入
	sm.Store("goku", 28)
	go func() {
		sm.Store("conan", 16)
	}()
	sm.Store("lufy", 23)

	// 2.读取
	age, _ := sm.Load("lufy")
	fmt.Println(age.(int))

	// 3.遍历
	sm.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})

	// 4.删除
	sm.Delete("goku")
	age, ok := sm.Load("goku")
	fmt.Println(age, ok)

	// 5.读取或写入
	sm.LoadOrStore("比克", 400)
	age, _ = sm.Load("比克")
	fmt.Println(age)
}
