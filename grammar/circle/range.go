package main

import "fmt"

/*
Go 不提供类似 C 支持的 while、do...while 等循环控制语法，而仅保留了一种语句，即 for 循环:

for i := 0; i < n; i++ {
 	//todo
}

但是,经典的三段式循环语句，需要获取迭代对象的长度 n,
鉴于此,为了更方便 Go 开发者对复合数据类型进行迭代,
例如 array、slice、channel、map,Go 提供了 for 循环的变体，即 for range 循环.

以下有几点需要注意：
*/

// 1.for range 中，参与循环表达式的只是对象的副本。
func demo1() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a) // [1 2 3 4 5]

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	/*
		r为什么不是 [1 12 13 4 5]？
		--	参与 for range 循环是 range 表达式的副本。也就是说，实际上参与循环的是 a 的副本（go 临时分配的连续字节序列，
		  	与 a 不是同一块内存空间），而不是真正的 a。
	*/
	fmt.Println("after for range loop, r =", r) // [1 2 3 4 5]
	fmt.Println("after for range loop, a =", a) //  [1 12 13 4 5]
}
func main() {
	demo1()
}
