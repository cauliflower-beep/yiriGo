package main

import (
	"testing"
)

// BenchmarkStr2Byte
//  @Description: go test -bench='Byte$'
//  @param b
func BenchmarkStr2Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte("Hello world")
	}
}

// BenchmarkStr2ByteOnce
//  @Description: go test -bench='Once$'
//  @param b
func BenchmarkStr2ByteOnce(b *testing.B) {
	data := []byte("Hello world")
	for i := 0; i < b.N; i++ {
		_ = data
	}
}
