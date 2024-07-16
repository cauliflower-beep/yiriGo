package main

import (
	"fmt"
	"sync"
)

/*
非安全模式
日常开发时，需要考虑如下问题：
1.对象创建时的线程安全问题；
2.是否支持延迟加载；
3.getInstance()性能是否够高
*/

type lazy struct{}

var instance_lazy *lazy

func GetLazy() *lazy {
	if instance_lazy == nil {
		fmt.Println("创建lazy实例")
		instance_lazy = &lazy{}
	}
	return instance_lazy
}

// 懒汉模式只有通过第一次调用GetLazy时才会创建,但上面的GetLazy在多线程环境下是不安全的，请看 singleTest.go 测试

/////////////////////////////////////////////////

/*
安全模式：
通过加锁的形式，实现并发环境下的线程安全。
*/

type lazy_safe struct{}

var instance_lazy_safe *lazy_safe

var m sync.Mutex

func GetLazySafe() *lazy_safe {
	m.Lock() // 加锁保证多线程安全
	defer m.Unlock()
	if instance_lazy_safe == nil {
		fmt.Println("创建安全lazy实例")
		instance_lazy_safe = &lazy_safe{}
	}
	return instance_lazy_safe
}

/*
上述写法中，每一次请求单例的时候，都会加锁和减锁
而锁的作用只在于解决对象初始化的时候可能出现的并发问题，当对象创建了之后，加锁就失去了意义，会拖慢速度
所以我们引入双重检查机制(check-lock-check)，也叫DCL(Double Check Lock),代码如下
*/

func GetLazyDcl() *lazy_safe {
	if instance_lazy_safe == nil {
		m.Lock()
		defer m.Unlock()
		if instance_lazy_safe == nil {
			instance_lazy_safe = &lazy_safe{}
		}
	}
	return instance_lazy_safe
}

//////////////////////////////////////////
/*
go专属安全模式
golang标准包中的once.Do函数，保证函数只被执行一次。
利用这个特性来实现单例，once.Do是多协程安全的
*/

type lazy_once struct{}

var instance_once *lazy_once

var once sync.Once

func GetLazyOnce() *lazy_once {
	once.Do(func() {
		fmt.Println("创建lazy_once实例")
		instance_once = &lazy_once{}
	})
	return instance_once
}

// 默写一个单例
