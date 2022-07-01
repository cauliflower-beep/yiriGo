package main

import (
	"io"
	"net/http"
)

/*
go语言中的错误处理一直被部分go开发者诟病，其中重复的 if err!=nil 样板代码严重降低代码的可读性。
这里介绍两种优化重复样板代码的方式
*/

/*
1.封装错误检查函数，代码中需要处理错误的地方，直接调用该函数。
	这种方式，虽然在视觉上提升了代码的可读性，但是需要特殊处理错误的场景也有局限性。
	比如需要使用额外的信息完善错误时，该方式并不适用此类场景。
*/
func CheckErrors(err error) {
	if err != nil {
		//do something
	}
}

/*
2.结构体中定义错误信息字段，将结构体的方法与错误信息绑定在一起
	阅读 Go 标准库 bufio 的代码片段，我们可以发现在 Writer 结构体中定义一个 err 字段，将错误信息封装在结构体中。
	在 Writer 结构体的方法的开头先判断 err 字段是否为 nil，如果 err 字段的值不是 nil，则直接返回 err，
	从而减少 if err != nil 样板代码的重复出现。
*/
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

func (w *Writer) Flush() error {
	if w.err != nil {
		return w.err
	}
	// ...
	return nil
}

func main() {
	_, err := http.Get("https://book.douban.com/")
	CheckErrors(err)
}

/*
总结
通过在结构体中定义错误信息的字段，将结构体的方法和错误信息绑定在一起的优化方式，相比较第一种方式更加优雅。
*/
