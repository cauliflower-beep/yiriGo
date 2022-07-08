package main

import "testing"

/*
 for range 使用的是副本数据，那 for range 会比经典的 for 循环消耗更多的资源并且性能更差吗？

基于副本复制问题，我们先使用基准示例来验证一下：对于大型数组，for range 是否一定比经典的 for 循环运行得慢？
	--从输出结果来看，for range 的确会稍劣于 for 循环，当然这其中包含了编译器级别优化的结果（通常是静态单赋值，或者 SSA 链接）。

让我们关闭优化开关，再次运行压力测试。（？）
https://mp.weixin.qq.com/s/XzEZm9xSzG_QEx_3QTpBnQ
	go test -c -gcflags '-N -l' . -o forRange1.test ./forRange1.test -test.bench .
*/

// 运行 go test .\range_test.go -bench . -benchmem 测试
func BenchmarkClassicForLoopIntArray(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]int
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr); j++ {
			arr[j] = j
		}
	}
}

func BenchmarkForRangeIntArray(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]int
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j] = j
			_ = v
		}
	}
}

/*
结论：
	1.不管是什么类型的结构体元素数组，经典的 for 循环遍历的性能比较一致，但是 for range 的遍历性能会随着结构字段数量的增加而降低。
	2.如果一个结构体类型有超过一定数量的字段（或一些其他条件），就会将该类型视为 unSSAable。如果 SSA 不可行，那么就无法通过 SSA 优化，
      这也是造成上面基准测试结果的重要原因。
	3.对于遍历大数组而言， for 循环能比 for range 循环更高效与稳定，这一点在数组元素为结构体类型更加明显。
	4.另外，由于在 Go 中切片的底层都是通过数组来存储数据，尽管有 for range 的副本复制问题，
	  但是切片副本指向的底层数组与原切片是一致的。这意味着，当我们将数组通过切片代替后，不管是通过 for range 或者 for 循环均能得到一致的稳定的遍历性能。
*/
