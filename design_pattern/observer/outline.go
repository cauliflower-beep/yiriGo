package main

import (
	"reflect"
	"sync"
)

/*
	观察者模式主要是用来实现事件驱动编程。
	事件驱动编程的应用还是挺广的，除了我们都知道的能够用来解耦：用户修改密码后，给用户发短信进行风险提示之类的典型场景，
	在微服务架构实现最终一致性、实现事件源（A + ES）这些都会用到。

	概念
		观察者模式 (Observer Pattern)，定义对象间的一种一对多依赖关系，
		使得每当一个对象状态发生改变时，其相关依赖对象皆得到通知.
		依赖对象在收到通知后，可自行调用自身的处理程序，实现想要干的事情，比如更新自己的状态。
		发布者对观察者唯一了解的是它实现了某个接口（观察者接口）。这种松散耦合的设计最大限度地减少了对象之间的相互依赖，因此使我们能够构建灵活的系统。

	理解
		1.观察者模式也经常被叫做发布 - 订阅（Publish/Subscribe）模式。
		  上面说的定义对象间的一种一对多依赖关系，一 - 指的是发布变更的主体对象，多 - 指的是订阅变更通知的订阅者对象。
		2.发布的状态变更信息会被包装到一个对象里，这个对象被称为事件，事件一般用英语过去式的语态来命名，
		  比如用户注册时，用户模块在用户创建好后发布一个事件 UserCreated 或者 UserWasCreated，这样从名字上就能看出，这是一个已经发生的事件。
		3.事件发布给订阅者的过程，其实就是遍历一下已经注册的事件订阅者，逐个去调用订阅者实现的观察者接口方法，
		  比如叫 handleEvent 之类的方法，这个方法的参数一般就是当前的事件对象。
		4.至于很多人会好奇的，事件的处理是不是异步的？
		  主要看我们的需求是什么，一般情况下是同步的，即发布事件后，触发事件的方法会阻塞等到全部订阅者返回后再继续，
		  当然也可以让订阅者的处理异步执行，完全看我们的需求。
		5.大部分场景下其实是同步执行的，单体架构会在一个数据库事务里持久化因为主体状态变更，而需要更改的所有实体类。
		6.微服务架构下常见的做法是有一个事件存储。
		  订阅者接到事件通知后，会把事件先存到事件存储里，这两步也需要在一个事务里完成才能保证最终一致性，
		  后面会再有其他线程把事件从事件存储里搞到消息设施里，发给其他服务，从而在微服务架构下实现各个位于不同服务的实体间的最终一致性。
		7.所以观察者模式，从程序效率上看，大多数情况下没啥提升，更多的是达到一种程序结构上的解耦，让代码不至于那么难维护。
*/

/*
	下面来实现一个简单的观察者模式
*/

//// Publish 接口，相当于是发布者的定义
//type Publish interface {
//	Subscribe(obs Observer)
//	Notify(msg string)
//}
//
//// 观察者接口
//type Observer interface {
//	Update(msg string)
//}
//
//// Publish 实现
//type PublishImp struct {
//	observers []Observer
//}
//
//// Subscribe 添加观察者（订阅者）
//func (pub *PublishImp) Subscribe(obs Observer) {
//	pub.observers = append(pub.observers, obs)
//}
//
//// Notify 发布通知
//func (pub *PublishImp) Notify(msg string) {
//	for _, obs := range pub.observers {
//		obs.Update(msg)
//	}
//}
//
//type Obs1 struct{}
//
//func (obs Obs1) Update(msg string) {
//	fmt.Println("im obs1 , have receive msg: ", msg)
//}
//
//type Obs2 struct{}
//
//func (obs Obs2) Update(msg string) {
//	fmt.Println("im obs2 , have receive msg: ", msg)
//}

/*
	以上就是 Go 实现观察者模式的代码
	实际应用的时候，一般会定义个事件总线 EventBus 或者事件分发器 Event Dispatcher，来管理事件和订阅者间的关系，以及分发事件，
	它们两个就是名不一样，角色定位一样。

	下面我们实现一个支持以下功能的事件总线：
	1.异步不阻塞
	2.支持任意参数值
	3.代码来自 https://lailin.xyz/post/observer.html
*/

// Bus
type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// AsyncEventBus 异步事件总线
type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

// NewAsyncEventBus new
func NewAsyncEventBus() *AsyncEventBus {
	return nil
}
func main() {
	// 观察者模式demo
	//// 注册订阅者
	//var pub PublishImp
	//pub.Subscribe(Obs1{})
	//pub.Subscribe(Obs2{})
	//pub.Notify("木叶飞舞之处，火亦生生不息")

	// 事件总线

}
