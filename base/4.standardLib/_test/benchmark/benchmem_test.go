package main

import (
	"math/rand"
	"testing"
	"time"
)

// go test -bench="Generate" -benchtime=5s -benchmem .

// generate
//  @Description: 生成一组长度为 n 的随机序列
func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// generateWithCap
//  @Description: 生成一组长度为 n 的随机序列
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	/*
		不过它在创建切片时，将切片的容量(capacity)设置为n，
		这样切片就会一次性申请 n 个整数所需的内存
	*/
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateWithCap(1000000)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(1000000)
	}
}
