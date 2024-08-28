package a_s

import (
	"fmt"
	"testing"
)

func TestNoAffect(t *testing.T) {
	var sl []int
	sl = append(sl, 1)
	fmt.Println("内部增加元素前，外部切片:", sl)
	noAffect(sl)
	fmt.Println("内部增加元素后，外部切片:", sl)
	// 可以看到，内部对切片的修改，外部不可见
}

func TestAffect(t *testing.T) {
	var sl []int
	sl = append(sl, 1)
	fmt.Println("before setData copy:", sl)
	affect(sl)
	fmt.Println("after setData copy:", sl)
	// 可以看到，内部对切片的修改，外部可见
}

func TestAddrSwitch(t *testing.T) {
	var sl []int
	sl = append(sl, 1)
	fmt.Printf("外部切片底层数组初始地址:%p|外部切片初始地址:%p\n", sl, &sl)
	slAddrSwitch(sl)
	fmt.Printf("外部切片底层数组扩容后地址:%p|外部切片扩容后地址:%p\n", sl, &sl)
	// 可以发现，内部扩容前后，内部切片指向的底层数组地址发生了变化，而外部没变，所以外部不可见
}

func TestSetData(t *testing.T) {
	setData()
}

func TestExpend(t *testing.T) {
	expand()
}
