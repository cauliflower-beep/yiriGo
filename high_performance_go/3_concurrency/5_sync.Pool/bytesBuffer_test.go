package main

import (
	"bytes"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

// go test -bench="Buffer*" . -benchmem
/*
	这个例子创建了一个bytes.Buffer 对象池，而且每次只执行一个简单的Write操作，纯粹的内存搬运工，耗时几乎可以忽略
	内存分配和回收的耗时占比较多，因此对整体程序的性能影响很大
*/
func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		/*
			加上 Reset() 操作，不会影响性能,且这里是必要的。
			如果不执行 Reset(),下次从pool中取出来再用时，就是在原来的 buf 上追加写了
		*/
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		// 这里 buf 每次都是新申请的 ，所以没有必要 reset()
		buf.Write(data)
	}
}
