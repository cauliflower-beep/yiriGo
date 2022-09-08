package main

import "fmt"

/*
	优化代码性能之前，首先要了解当前性能怎么样。
	Go 语言标准库内置的 testing 测试框架提供了基准测试(benchmark)的能力，能让我们很容易地对某一段代码进行性能测试
*/

// 计算第 N 个斐波那契数

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fmt.Println(fib(20))
}
