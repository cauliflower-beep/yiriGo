package main

import "testing"

// 实现一个 benchmark 用例

/*
	benchmark 和普通的单元测试用例一样，都位于 _test.go 文件中;
	函数名以 Benchmark 开头，参数是 b *testing.B;
	和普通的单元测试用例很像，单元测试函数名以 Test 开头，参数是 t *testing.T
*/

/*
	运行当前 package 内的用例：go test example 或 go test .
	运行子 package 内的用例： go test example/<package name> 或 go test ./<package name>
	如果想递归测试当前目录下的所有的 package：go test ./... 或 go test example/...

	go test 命令默认不运行 benchmark 用例的，如果我们想运行 benchmark 用例，则需要加上 -bench 参数:
		go test -bench .
	-bench 参数支持传入一个正则表达式，匹配到的用例才会得到执行，例如，只运行以 Fib 结尾的 benchmark 用例：
		go test -bench='Fib$' .

	结果分析可参考博客：
	https://geektutu.com/post/hpg-benchmark.html

	benchmark常用参数：
		• go test -bench . 执行当前测试
		• b.N 决定用例需要执行的次数
		• -bench 可传入正则，匹配用例
		• -cpu 传入列表，可改变 CPU 核数。测试并发性能的时候可用，如果一个测试是单线程，那改cpu核数意义不大
		• -benchtime 可指定执行时间或具体次数
		• -count 可设置 benchmark 轮数
		• -benchmem 可查看内存分配量和分配次数
*/
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}
