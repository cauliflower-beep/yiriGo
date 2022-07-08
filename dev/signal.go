package main

/*
我们在程序的处理中，可能会有一些场景，如：
程序被退出时，将内存中的数据输入文件或db，或者其他必要的处理

针对特定系统信号针对性处理操作，如可以通过发送信号进行数据采集、日志输出、配置加载等操作,
这些情况下就需要针对Signal的处理了，Go中关于信号的处理主要集中于os/signal package中。

os/signal中涉及的function主要有：Notify、Stop、Ignore、Reset、NotifyContext。
在了解function之前，我们先来看看信号是如何存储的，os/signal中信号的是存储在handlers。
*/

// handlers
/*
var handlers struct {
	sync.Mutex
	// Map a channel to the signals that should be sent to it.
	m map[chan<- os.Signal]*handler
	// Map a signal to the number of channels receiving it.
	ref [numSig]int64
	// Map channels to signals while the channel is being stopped.
	// Not a map because entries live here only very briefly.
	// We need a separate container because we need m to correspond to ref
	// at all times, and we also need to keep track of the *handler
	// value for a channel being stopped. See the Stop function.
	stopping []stopping
}

type stopping struct {
	c chan<- os.Signal
	h *handler
}

type handler struct {
	mask [(numSig + 31) / 32]uint32 // mask是一个长度为3的uint32数组，意味着1个handler可以包含96个信号
}

func (h *handler) want(sig int) bool {
	return (h.mask[sig/32]>>uint(sig&31))&1 != 0 // 和存储相反，右移获取指定位是否为1
}

func (h *handler) set(sig int) {
	h.mask[sig/32] |= 1 << uint(sig&31) // 存储时，通过sig/32获取大致存储位置，同时左移sig&31作为具体存储位置。
}

func (h *handler) clear(sig int) { // 通过异或清零指定位
	h.mask[sig/32] &^= 1 << uint(sig&31)
}
*/
