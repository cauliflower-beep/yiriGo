package int

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStr2Runes(t *testing.T) {

	s := "国"
	fmt.Println(len(s))
	fmt.Println([]rune(s))

	// golang中的单个字符都用 rune 类型表示，不区分中英文
	b := 'b'
	fmt.Println("单字符 b 的类型：", reflect.TypeOf(b)) // int32
	r := '中'
	fmt.Println(reflect.TypeOf(r)) // int32
}

func TestGetCharLenByRuneArr(t *testing.T) {
	s := "我的玫瑰，种在繁星中的一颗"
	fmt.Println("直接统计：", len(s))

	bs := []byte(s)
	fmt.Println("通过byte数组统计：", len(bs))

	// 转成[]rune之后，字符串中的每个字符无论占多少个字节都用int32来表示，所以可以正确处理中文
	rs := []rune(s)
	fmt.Println("通过rune数组统计：", len(rs))
}
