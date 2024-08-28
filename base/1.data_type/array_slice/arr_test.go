package a_s

import (
	"fmt"
	"testing"
)

func TestDefine(t *testing.T) {
	// 数组定义
	define()
}

func TestTransfer(t *testing.T) {
	// 数组是值传递
	var arr1, arr2 [5]int
	setDataCopy(arr1)
	setDataAddr(&arr2)
	fmt.Printf("l1:%v\nl2:%v\n", arr1, arr2)
}

func TestArrAddr(t *testing.T) {
	arrAddr()
}
