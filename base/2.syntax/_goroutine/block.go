package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var v atomic.Value

var mu sync.Mutex

func Done() <-chan struct{} {
	fmt.Println("here.")
	d := v.Load()
	if d != nil {
		return d.(chan struct{})
	}
	mu.Lock()
	defer mu.Unlock()
	d = v.Load()
	if d == nil {
		d = make(chan struct{})
		v.Store(d)
	}
	return d.(chan struct{})
}

func taskX() {
	for {
		select {
		case <-Done(): // 通道关闭||通道赋值都可以解除这里的阻塞
			fmt.Println("task over.")
			return
		default:
			fmt.Println("task running.")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	go taskX()

	time.Sleep(time.Second * 5)
	d, _ := v.Load().(chan struct{})

	//close(d)
	//fmt.Println("close chan.")

	d <- struct{}{}
	fmt.Println("assign chan.")

	time.Sleep(time.Second * 5)
	fmt.Println("main over.")
}
