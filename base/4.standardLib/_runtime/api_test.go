package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestCPU(t *testing.T) {
	// 操作系统/CPU型号
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
