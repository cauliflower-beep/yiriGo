package main

import "testing"

/*
	不同的函数复杂度不同，O(1)，O(n)，O(n^2) 等，利用 benchmark 验证复杂度一个简单的方式，是构造不同的输入。
	仍以 generate 内存测试为例，稍加改造

	运行测试：
		go test -bench="0$" .
	结果分析：
		goos: windows
		goarch: amd64
		pkg: yiriGo/high_performance_go/1_performanceAnalysis/1_benchmark
		cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
		BenchmarkGenerate1000-12           45516             26033 ns/op
		BenchmarkGenerate10000-12           6000            200992 ns/op
		BenchmarkGenerate100000-12           588           1999295 ns/op
		BenchmarkGenerate1000000-12           60          20085595 ns/op
		PASS
		ok      yiriGo/high_performance_go/1_performanceAnalysis/1_benchmark    5.486s
	可以发现复杂度是线性的：输入变为原来的10倍左右，函数每次调用时长也差不多是原来的10倍
*/

func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(i)
	}
}

// 构造 4 个不同输入的benchmark 用例
func BenchmarkGenerate1000(b *testing.B)    { benchmarkGenerate(1000, b) }
func BenchmarkGenerate10000(b *testing.B)   { benchmarkGenerate(10000, b) }
func BenchmarkGenerate100000(b *testing.B)  { benchmarkGenerate(100000, b) }
func BenchmarkGenerate1000000(b *testing.B) { benchmarkGenerate(1000000, b) }
