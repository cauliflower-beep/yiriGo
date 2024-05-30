package main

import "testing"

/*
	benchmark常用参数：
		• go test -bench . 执行当前测试
		• b.N 决定用例需要执行的次数
		• -bench 可传入正则，匹配用例
		• -cpu 传入列表，可改变 CPU 核数。测试并发性能的时候可用，如果一个测试是单线程，那改cpu核数意义不大
		• -benchtime 可指定执行时间或具体次数
		• -count 可设置 benchmark 轮数
		• -benchmem 可查看内存分配量和分配次数
*/
/*
	benchmark 用例和普通的单元测试用例一样，都位于 _test.go 文件中;
	函数名以 Benchmark 开头，参数是 b *testing.B;
*/

// BenchmarkFib
//  @Description: go test -bench='Fib$' .  默认时间是1s
// 	@Description: go test -bench='Fib$' -benchtime=5s .  指定运行时间为5s，增加测试次数以提升测试准确度
// 	@Description: go test -bench='Fib$' -benchtime=30x .  指定运行时间为30次
// 	@Description: go test -bench='Fib$' -benchtime=5s -count=3 .  指定执行3轮
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}
