package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
关键词：golang slice两个冒号的理解
*/

func main() {
	// 所有空切片指向的地址都是一致的
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	// 获取切片底层对应的数据结构
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s1)), "\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))

	// nil切片的地址是空的
	var s3 []int
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s3)))
}
