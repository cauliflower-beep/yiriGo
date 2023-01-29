package main

import (
	"bytes"
	"fmt"
)

/*
	文章链接：http://events.jianshu.io/p/cb12e88c60d6
	https://blog.csdn.net/u013007900/article/details/89126811
	io包中最重要的两个接口就是 Reader 和 Writer 只要实现了这两个接口，它就有了 io 的能力
	io.Reader 的原型：
		type Reader interface {
			// 将内容读到 []byte 中去，返回读到的字节数以及读取过程中的错误
			Read(p []byte) (n int, err error)
		}
	io.Writer 跟 io.Reader 一样，都是 Interface 类型，功能非常强大，在任何需要写入数据，处理数据流的地方，我们都应该尽可能使用这两个类型的对象。
	io.Writer 的原型：
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	跟 io.Reader 类似，一个对象只要实现了 Write() 函数，这个对象就自动成为 Writer 类型。
	常见 Writer 类型
	（1）文件操作
		使用 os.Create() 创建文件时，会返回一个 os.File 对象，它是一个 struct，但是由于它实现了 Read() ，Write()，Closer() 等函数，
		因此它同时也是 Reader, Writer, Closer 等类型。
	（2）bytes.Buffer
		在 Go 语言中，string 类型是 immutable 的，因此它没有对应的 Writer，也就是说不存在 strings.NewWriter(s) 这种函数。
		最好的替代方式就是使用 bytes.Buffer，因为它既是一个 Reader 也是一个 Writer，我们既可以往里面写也可以往外读。
		我们可以通过 buf.String() 得到 string 类型的数据，也可以通过 buf.Bytes() 拿到 []byte 类型的数据。
		下面的例子展示了我们通过 bytes.NewBufferString(s) 函数，先在 buffer 中初始化一段 string，然后往里面 append 另外一段 string。
	（3）http.ResponseWriter
		在使用 Go 语言进行 Web 开发时，http.ResponseWriter 是最基本的类型之一，它本身是一个 Interface 类，原型可自行查看源码。
		它只申明了需要实现三个函数，由于其要求了 Writer() 函数，包含了 Writer 的要求，因此，任何是符合 ResponserWriter 的类型必然是 Writer 类型。
*/

// buffer
//  @Description: （2）bytes.Buffer 案例
func buffer() {
	s := "Hello"
	buf := bytes.NewBufferString(s)
	s2 := "to be appended"
	buf.WriteString(s2) // 或者 fmt.Fprint(buf, s2)
	fmt.Println("Final string:", buf.String())
	fmt.Println("Final buffer:", buf.Bytes())
}
func main() {
	buffer()
}
