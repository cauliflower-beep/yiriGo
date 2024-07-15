package main

import "sync"

type T int

// safeClose 简单粗暴的方案关闭可能关闭的通道
func safeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// the return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch) // _panic if ch is closed
	return true
}

// safeSend 简单粗暴的方案向一个可能关闭的channel中发送数据
func safeSend(ch chan T, val T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- val // _panic if ch is closed
	return false
}

// MyChannel 礼貌的关闭channel—sync.Once 版
type MyChannel struct {
	C    chan int
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan int)}
}

func (mc *MyChannel) SafeClose() {
	// 礼貌的关闭channel
	mc.once.Do(func() {
		close(mc.C)
	})
}
