package main

import "fmt"

// fib
//  @Description: 计算第 N 个斐波那契数
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
