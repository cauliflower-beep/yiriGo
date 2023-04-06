package main

import "fmt"

// https://juejin.cn/post/7195388141115539493#heading-6

// printSlice[T any] [T any]参数的类型，意思是该函数支持任何T类型;
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}

func main() {
	printSlice[int]([]int{66, 77, 88, 99, 100})
	printSlice[float64]([]float64{1.1, 2.2, 5.5})
	printSlice[string]([]string{"烤鸡", "烤鸭", "烤鱼", "烤面筋"})
	//省略显示类型
	printSlice([]int64{55, 44, 33, 22, 11})
}

/*
 	多个泛型参数语法：
	[T, M any]
	[T any, M any]
	[T any, M comparable]
*/
