package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// 测试 strings 中的replace函数
func replaceTest() {
	old := "d"
	new := "D"
	fmt.Println(strings.Replace("aabbccddd", old, new, -1))
}

// 测试 uuid 的生成
func uuidTest() {
	for i := 0; i < 5; i++ {
		fmt.Println(uuid.New().String())
	}
}
func main() {
	replaceTest()

	uuidTest()
}
