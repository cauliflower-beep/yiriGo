package main

/*
带有类型参数的类型被叫做泛型类型。
下面定义一个底层类型为切片类型的新类型 vector,它可以存储任何类型的的切片
要使用泛型类型，要先对其进行实例化，就是给类型参数指定一个实参。
*/
type vector[T any] []T

func main() {
	v := vector[int]{58, 1881}
	printSlice(v)
	v2 := vector[string]{"烤鸡", "烤鸭", "烤鱼", "烤面筋"}
	printSlice(v2)
}
