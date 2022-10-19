package gorm

import "fmt"

// Println 封装一层，所有的打印都走这里
func Println(args ...interface{}) {
	fmt.Println(args...) // 不需要打印就注掉这一行
}
