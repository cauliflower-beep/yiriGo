package int

import "fmt"

func newInt() {
	a := new(int)
	fmt.Printf("a指向的地址：%p|a的值：%d\n", a, *a)
	*a = 1
	fmt.Printf("a指向的地址：%p|a的值：%d\n", a, *a)

	var b int
	fmt.Printf("b变量的地址：%p|b的值：%d\n", &b, b)
	b = 1
	fmt.Printf("b变量的地址：%p|b的值：%d\n", &b, b)
}
