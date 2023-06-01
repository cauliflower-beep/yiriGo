package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

/*
	字符串高效拼接
	在 Golang 中，string 是不可变的。拼接字符串事实上是创建了一个新的字符串对象
	如果代码中存在大量的字符串拼接，对性能会产生严重的影响。
*/

// 1.1 常见的拼接方式
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// randomString
// @Description: 为了避免编译器优化，我们首先实现一个函数，可以生成长度为 n 的随机字符串。
// 然后利用这个函数生成字符串 str，将 str 拼接 N 次
func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 在 Go 语言中，常见的字符串拼接方式有 5 种 (漏了一种strings.join(s,"."))

// plusConcat
// @Description: 使用 +
func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

// sprintfConcat
// @Description:使用 fmt.Sprintf
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

// builderConcat
// @Description:strings.Builder
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// bufferConcat
// @Description:使用 bytes.Buffer
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

// byteConcat
// @Description:使用 []byte
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

// preByteConcat
// @Description:如果长度是可预知的，那么创建 []byte 时，我们还可以预分配切片的容量(cap)。
func preByteConcat(n int, str string) string {
	// make([]byte, 0, n*len(str)) 第二个参数是长度，第三个参数是容量(cap)，切片创建时，将预分配 cap 大小的内存
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

/*
	1.2 benchmarkmark 性能比拼
	运行 stringConcat_test.go
	从基准测试的结果来看，使用 + 和 fmt.Sprintf 的效率是最低的，和其余的方式相比，性能相差约 1000 倍，而且消耗了超过 1000 倍的内存。
	当然 fmt.Sprintf 通常是用来格式化字符串的，一般不会用来拼接字符串。
	strings.Builder、bytes.Buffer 和 []byte 的性能差距不大，而且消耗的内存也十分接近，性能最好且消耗内存最小的是 preByteConcat，
	这种方式预分配了内存，在字符串拼接的过程中，不需要进行字符串的拷贝，也不需要分配新的内存，因此性能最好，且内存消耗最小。
*/

/*
	1.3 建议
	综合易用性和性能，一般推荐使用 strings.Builder 来拼接字符串。
	这是 Go 官方对 strings.Builder 的解释：
	A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.
*/
// preBuilderConcat
// @Description:string.Builder 也提供了预分配内存的方式 Grow
func preBuilderConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// 使用了 Grow 优化后的版本的 benchmark 结果如下：(自行验证)
// 与预分配内存的 []byte 相比，因为省去了 []byte 和字符串(string) 之间的转换，内存分配次数还减少了 1 次，内存消耗减半。

/*
	2 性能背后的原理
	2.1 比较 strings.Builder 和 +
	strings.Builder 和 + 性能和内存消耗差距如此巨大，是因为两者的内存分配方式不一样。
	字符串在 Go 语言中是不可变类型，占用内存大小是固定的。
	当使用 + 拼接 2 个字符串时，生成一个新的字符串，那么就需要开辟一段新的空间，新空间的大小是原来两个字符串的大小之和。
	拼接第三个字符串时，再开辟一段新空间，新空间大小是三个字符串大小之和，以此类推。
	假设一个字符串大小为 10 byte，拼接 1w 次，需要申请的内存大小为：
		10 + 2 * 10 + 3 * 10 + ... + 10000 * 10 byte = 500 MB
	而 strings.Builder，bytes.Buffer，包括切片 []byte 的内存是以倍数申请的。
	例如，初始大小为 0，当第一次写入大小为 10 byte 的字符串时，则会申请大小为 16 byte 的内存（恰好大于 10 byte 的 2 的指数），
	第二次写入 10 byte 时，内存不够，则申请 32 byte 的内存，第三次写入内存足够，则不申请新的，以此类推。
	在实际过程中，超过一定大小，比如 2048 byte 后，申请策略上会有些许调整。
	我们可以通过打印 builder.Cap() 查看字符串拼接过程中，strings.Builder 的内存申请过程，见 stringConcat_test.go/TestBuilderConcat 测试用例

	可以看到，2048 以前按倍数申请，2048 之后，以 640 递增，最后一次递增 24576 到 122880。总共申请的内存大小约 0.52 MB，约为上一种方式的千分之一。
		16 + 32 + 64 + ... + 122880 = 0.52 MB

	2.2 比较 strings.Builder 和 bytes.Buffer
	strings.Builder 和 bytes.Buffer 底层都是 []byte 数组，但 strings.Builder 性能比 bytes.Buffer 略快约 10% 。
	一个比较重要的区别在于，bytes.Buffer 转化为字符串时重新申请了一块空间，存放生成的字符串变量，而 strings.Builder 直接将底层的 []byte 转换成了字符串类型返回了回来。

	bytes.Buffer
	// To build strings more efficiently, see the strings.Builder type.
	func (b *Buffer) String() string {
		if b == nil {
			// Special case, useful in debugging.
			return "<nil>"
		}
		return string(b.buf[b.off:])
	}

	strings.Builder
	// String returns the accumulated string.
	func (b *Builder) String() string {
		return *(*string)(unsafe.Pointer(&b.buf))
	}

	bytes.Buffer 的注释中还特意提到了：
	To build strings more efficiently, see the strings.Builder type.
*/
