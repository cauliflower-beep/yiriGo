package main

import (
	"fmt"
	"runtime"
)

func Add(a, b int) int {
	_, file, line, ok := runtime.Caller(1) // 获取调用者文件及代码堆栈信息，打印日志时非常有用
	fmt.Printf("file:%s|line:%d|ok:%v\n", file, line, ok)
	return a + b
}
