package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// BenchmarkFmt
//  @Description: go test -bench='Fmt$' 运行基准测试
//  @param b
func BenchmarkFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(rand.Int())
	}
}

// BenchmarkStrconv
//  @Description: go test -bench='Strconv$'
//  @param b
func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(rand.Int())
	}
}
