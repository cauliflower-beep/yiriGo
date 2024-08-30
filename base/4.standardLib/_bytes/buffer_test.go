package main

import (
	"bytes"
	"testing"
)

// TestWrite 测试写入
func TestWrite(t *testing.T) {
	var b bytes.Buffer
	for i := 0; i < 10; i++ {
		b.WriteString("hello")
	}
	t.Log(b.String())
}