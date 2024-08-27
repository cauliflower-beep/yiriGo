package int

import (
	"fmt"
	"testing"
)

func TestChar2Byte(t *testing.T) {
	/*
		`中`是rune(int32)类型，转换为byte(uint8)会导致截断
		`中`的Unicode编码为 4E2D
		转换为二进制就是 100111000101101
		截断低8位转换为10进制表示：00101101->45

	*/
	c := '中'
	fmt.Println(byte(c)) // 45
}

func TestStr2Bytes(t *testing.T) {
	s := "国"
	fmt.Println([]byte(s))
}
