package main

import "fmt"

/*
	到这一章的时候你说巧不巧，
	下午同事刚在 tarsGo 的源码中看到这种写法，问我啥意思，当时我肯定不懂了
	但看到这一节就清楚了哈哈
*/
type phone interface {
	Sms()
}

// Iphone 为实现phone接口的导出类型 需将其作为接口api的一部分进行检查
type Iphone struct{}

type xiaomi struct{}

// 赋值的后面是类型断言的零值
var _ phone = Iphone{}
var _ phone = (*xiaomi)(nil)

func (x *xiaomi) Sms() {
	fmt.Println("身为小米，能发短信不是很正常?")
}
func (i Iphone) Sms() {
	fmt.Println("身为一个手机，能发短信还不是很正常?")
}

func main() {
	var iphone13 Iphone
	iphone13.Sms()
}
