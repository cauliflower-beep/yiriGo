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
	var l1, l2 []int
	fillNumCopy(l1)
	fillNumAddr(&l2)
	fmt.Printf("l1:%v\nl2:%v\n", l1, l2)
}

func TestArrAddr(t *testing.T) {
	arrAddr()
}
