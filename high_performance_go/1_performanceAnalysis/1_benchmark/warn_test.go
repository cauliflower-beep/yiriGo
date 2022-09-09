package main

import (
	"testing"
	"time"
)

/*
	如果在 benchmark 开始前，需要一些准备工作，
	并且准备工作比较耗时，则需要将这部分代码的耗时忽略掉。
*/
// BenchmarkFib2 运行测试 go test -bench='Fib2' -benchtime=50x .
func BenchmarkFib2(b *testing.B) {
	time.Sleep(time.Second * 3) // 模拟耗时准备任务
	// 如果不重置定时器，会发现每次调用几乎是重置定时器的 10 倍
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(30) // run fib(30) b.N times
	}
}

/*
	还有一种情况，如果每次函数调用前后需要一些准备工作和清理工作，
	我们可以使用 StopTimer 暂停计时以及使用 StartTimer 开始计时。

	例如，如果测试一个冒泡函数的性能，每次调用冒泡函数前，需要随机生成一个数字序列，这是非常耗时的操作
	这种场景下，就需要使用 StopTimer 和 StartTimer 避免将这部分时间计算在内。
*/
func bubbleSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] > nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nil
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		nums := generateWithCap(10000)
		b.StartTimer()
		bubbleSort(nums)
	}
}
