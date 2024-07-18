package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out") // 创建trace文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f) // 启动 trace goroutine
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("hello GMP!")
}
