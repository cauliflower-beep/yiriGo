package _int

import (
	"reflect"
)

/*
rune占用4个字节，即32个比特位，表示的是 Unicode 表中的一个字符
Unicode是一个可以表示世界范围内的绝大部分字符的编码规范
rune实质上就是int32类型，它用来表示Unicode的code point
由于 byte 能表示的值有限，只有2^8 = 256个，所以如果想表示中文的话，只能使用rune类型

规范：
rune alias for int32
*/

// type of param
func typeOfparam(s interface{}) reflect.Type {
	// fmt.Printf("%t",s)
	t := reflect.TypeOf(s)
	// fmt.Println(t)
	return t
}

// string->[]rune
func str2rs(s string) []rune {
	rs := []rune(s)
	return rs
}

// []rune->string
func rs2str(rs []rune) string {
	s := string(rs)
	return s
}

// func main() {
// 	// type compare
// 	r := []rune{2, 5}
// 	i := []int32{34, 67}
// 	fmt.Println(typeOfparam(r) == typeOfparam(i)) // true
//
// 	// convert
// 	fmt.Println(str2rs("中国人"))
// 	fmt.Println(rs2str([]rune{20013, 22269, 20154}))
//
// }

/*
https://juejin.cn/post/6931523990280208392

https://blog.csdn.net/joeyoj/article/details/135723403
*/
