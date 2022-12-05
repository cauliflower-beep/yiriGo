package main

import "fmt"

/*
	1. 数组定义的几种方式
*/
func define() {
	// 1.默认初始方式
	var names [3]string
	fmt.Println(names)

	// 2.初始化数组并赋值，{}中的元素个数不能大于[]中的数字
	var balance = [5]int{1, 2, 3}
	fmt.Println(balance)

	// 3.函数内初始化,若数组长度位置出现"..."，表示数组的长度是根据初始化值的个数计算的
	q := [...]int{1, 2, 3}
	fmt.Println(q)

	// 4.下面定义一个数组，长度为100，前99项都是0，最后一项是1
	r := [...]int{9: 1}
	fmt.Println(r)
	m := [...]int{0: 1, 2: 2, 5: 3}
	fmt.Println(m)

	// 6.指针数组
	x, y := 1, 2
	a := [...]*int{&x, &y}
	fmt.Println(a)
}

/*
	2.数组是值传递
*/
func fillNumCopy(l []int) {
	l = append(l, 1)
}
func fillNumAddr(l *[]int) {
	*l = append(*l, 2)
}
func main() {
	define()
	// 测试值传递还是地址传递
	var l1, l2 []int
	fillNumCopy(l1)
	fillNumAddr(&l2)
	fmt.Println("l1:", l1, "\n", "l2:", l2)
}
