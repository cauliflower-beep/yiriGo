package int

import (
	"fmt"
	"reflect"
)

/*
byte占用一个字节，即8个比特位，表示的是 ASCII 表中的一个字符
byte实质上就是uint8类型，只不过它强调数据是原始数据(raw data);而非数字
规范：
byte  alias for uint8

在golang中经常会用到[]byte和string的相互转化，尤其是在使用 json.Marshal 和 json.Unmarshal 的时候
*/

/*
为什么要多出一个 byte(rune) 类型？
理由很简单，uint8(int32),直观上会让人以为这是一个数值，但实际上也可以表示一个字符。为了消除这种直观错觉，就诞生了 byte(rune) 类型
*/

/*
补充：
go语言的string是用utf8进行编码的，英文字母占用一个字节，中文字符占用3个字节
以 “中” 字为例，utf8的10进制编码为 14989485，
转换为 2进制为 11100100 10111000 10101101
对应的三个字节为：228 184 173
*/
// type of slice
func typeOfSlice(s interface{}) reflect.Type {
	// fmt.Printf("%t",s)
	return reflect.TypeOf(s)
}

// string->[]byte
func str2bs(s string) []byte {
	bs := []byte(s)
	return bs
}

// []byte->string
func bs2str(bs []byte) string {
	s := string(bs)
	return s
}
func main() {
	// type of byte
	var a byte = 'a'               // 注意单双引号的区别，golang用单引号定义字符，类型为 uint8;用双引号定义字符串，类型为string
	fmt.Println(reflect.TypeOf(a)) // uint8

	// type compare
	b := []byte{2, 5}
	u := []uint8{34, 67}
	fmt.Println(typeOfSlice(b) == typeOfSlice(u)) // true

	// convert
	// fmt.Println(str2bs("abc"))
	fmt.Println(str2bs("中国人")) // [228 184 173 229 155 189 228 186 186]
	fmt.Println(bs2str([]byte{97, 98, 99}))

}

/*
https://juejin.cn/post/6931523990280208392
*/
