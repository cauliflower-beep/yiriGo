package main

import "fmt"

/*
	数组定义的几种方式:
*/
func define() {
	var names [3]string
	fmt.Println("1.默认初始化方式|", names)

	// 数组长度在定义之后无法再次修改 {}中的元素个数不能大于[]中定义的数组长度
	var balance = [5]int{1, 2, 3}
	fmt.Println("2.初始化数组并赋值|", balance)

	// 若数组长度位置出现"..."，表示数组的长度是根据初始化值的个数计算的
	q := [...]int{1, 2, 3}
	fmt.Println("3.函数内初始化|", q)

	// 下面定义一个数组，长度为10，前9项都是0，最后一项是1
	r := [...]int{9: 1}
	m := [...]int{0: 1, 2: 2, 5: 3}
	fmt.Printf("4.特殊方式初始化|r:%v^length:%d|%v\n", r, len(r), m)

	x, y := 1, 2
	a := [...]*int{&x, &y}
	fmt.Println("5.指针数组初始化|", a)
}

func main() {
	define()
}
