package _int

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

func TestSizeofInt(t *testing.T) {
	var i1 = 1
	var i2 int8 = -2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	// unsafe.Sizeof() 只返回数据类型的大小，不管引用数据的大小，单位时Byte
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))
	fmt.Println(unsafe.Sizeof(i3))
	fmt.Println(unsafe.Sizeof(i4))
	fmt.Println(unsafe.Sizeof(i5))

	// int的长度
	fmt.Println(strconv.IntSize)
}
