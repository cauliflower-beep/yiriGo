package main

import "fmt"

/*
	2.数组是值类型，每次传递都将产生一分副本
*/
func fillNumCopy(l []int) {
	l = append(l, 1)
}
func fillNumAddr(l *[]int) {
	*l = append(*l, 2)
}
func main() {
	var l1, l2 []int
	fillNumCopy(l1)
	fillNumAddr(&l2)
	fmt.Printf("l1:%v\nl2:%v", l1, l2)
}
