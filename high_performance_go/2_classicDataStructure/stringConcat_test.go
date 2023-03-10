package main

import (
	"fmt"
	"strings"
	"testing"
)

func benchmark(b *testing.B, f func(int, string) string) {
	// 每个 benchmark 用例中，生成了一个长度为 10 的字符串，并拼接 1w 次。
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

// 运行该用例 go test -bench="Concat$" -benchmem .
func BenchmarkPlusConcat(b *testing.B)    { benchmark(b, plusConcat) }
func BenchmarkSprintfConcat(b *testing.B) { benchmark(b, sprintfConcat) }
func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, builderConcat) }
func BenchmarkBufferConcat(b *testing.B)  { benchmark(b, bufferConcat) }
func BenchmarkByteConcat(b *testing.B)    { benchmark(b, byteConcat) }
func BenchmarkPreByteConcat(b *testing.B) { benchmark(b, preByteConcat) }

// TestBuilderConcat
//  @Description:go test -run="TestBuilderConcat" . -v
func TestBuilderConcat(t *testing.T) {
	var str = randomString(10)
	var builder strings.Builder
	cap := 0
	for i := 0; i < 10000; i++ {
		if builder.Cap() != cap {
			fmt.Print(builder.Cap(), " ")
			cap = builder.Cap()
		}
		builder.WriteString(str)
	}
}
