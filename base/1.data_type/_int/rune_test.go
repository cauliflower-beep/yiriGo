package _int

import (
	"fmt"
	"testing"
)

func TestGetCharLenByRuneArr(t *testing.T) {
	s := "我的玫瑰，种在繁星中的一颗"
	fmt.Println("直接统计：", len(s))
	// 转成[]rune之后，字符串中的每个字符无论占多少个字节都用int32来表示，所以可以正确处理中文
	rs := []rune(s)
	fmt.Println("通过rune数组统计：", len(rs))
}
