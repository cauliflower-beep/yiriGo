package main

import (
	"runtime"
	"testing"
	"time"
)

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		// 调用1000次timeout，即启动1000个子协程
		_ = timeout(f)
	}
	time.Sleep(time.Second * 2)   // 确保f执行完毕
	t.Log(runtime.NumGoroutine()) // 打印当前程序的协程个数
}

// TestBadTimeout
//  @Description: go test -run ^TestBadTimeout$ . -v
//  @param t
func TestBadTimeout(t *testing.T) {
	/*
		最终程序中存在着 1002 个子协程，说明即使是函数执行完成，协程也没有正常退出。
		如果在实际业务中，我们使用了上述代码，越来越多的协程会残留在程序中，最终导致内存耗尽(每个协程约占 2K 空间)，程序崩溃
	*/
	test(t, doBadthing)
}

// TestTimeoutWithBuffer
//  @Description: go test -run ^TestTimeoutWithBuffer$ . -v
//  @param t
func TestTimeoutWithBuffer(t *testing.T) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		_ = timeoutWithBuffer(doBadthing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
	// 协程数下降为2 创建的1000个子协程成功退出
}

// TestGoodTimeout
//  @Description: go test -run ^TestGoodTimeout$ . -v
func TestGoodTimeout(t *testing.T) {
	test(t, doGoodthing)
	// 协程数下降为2 创建的1000个子协程成功退出
}

// Test2phasesTimeout
//  @Description: go test -run ^Test2phasesTimeout$ . -v
func Test2phasesTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_ = timeoutFirstPhase()
	}
	time.Sleep(time.Second * 3)
	t.Log(runtime.NumGoroutine())
}
