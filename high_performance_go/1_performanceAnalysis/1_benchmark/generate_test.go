package main

import (
	"math/rand"
	"testing"
	"time"
)

/*
	内存分配次数也和性能是息息相关的。
	例如不合理的切片容量，将导致内存重新分配，带来不必要的开销。

	-benchmem 参数可以度量内存分配的次数:
		go test -bench="Generate" -benchtime=5s -benchmem .
	结果分析：
		goos: windows
		goarch: amd64
		pkg: yiriGo/high_performance_go/1_performanceAnalysis/1_benchmark
		cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
		BenchmarkGenerateWithCap-12          375          15918206 ns/op         8003590 B/op          1 allocs/op
		BenchmarkGenerate-12                 298          19981736 ns/op        45188467 B/op         41 allocs/op
		PASS
		ok      yiriGo/high_performance_go/1_performanceAnalysis/1_benchmark    15.602s
	可以看到，generate 分配的内存是 generateWithCap 的6倍，不设置切片容量的话，内存分配了 40 次
*/

// generate 生成一组长度为 n 的随机序列
func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

/*
	generateWithCap 作用同 generate；
	不过它在创建切片时，将切片的容量(capacity)设置为n，
	这样切片就会一次性申请 n 个整数所需的内存
*/
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
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
