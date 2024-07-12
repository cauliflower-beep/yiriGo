package main

import "fmt"

func main() {
	const n1 = 1
	fmt.Println(n1, "|", &n1) // 编译失败，获取不到n1地址
}
