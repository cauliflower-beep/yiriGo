package main

import "testing"

/*指定Map容量提示*/
// BenchmarkMapWithoutHint
//  @Description: go test -bench='WithoutHint'
//  @param b
func BenchmarkMapWithoutHint(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

// BenchmarkMapWithHint
//  @Description: go test -bench='WithHint'
//  @param b
func BenchmarkMapWithHint(b *testing.B) {
	m := make(map[int]int, b.N)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

/*指定slice容量*/
// BenchmarkSliceWithoutCap
//  @Description: go test -bench='WithoutCap'
//  @param b
func BenchmarkSliceWithoutCap(b *testing.B) {
	s := make([]int, 0)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}

// BenchmarkSliceWithCap
//  @Description: go test -bench='WithCap'
//  @param b
func BenchmarkSliceWithCap(b *testing.B) {
	s := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}
