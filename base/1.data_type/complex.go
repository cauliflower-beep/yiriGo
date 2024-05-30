package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := complex(1, 2)
	b := []int{4, 5, 6}
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(real(a)), imag(a))
	fmt.Println(b[real(2)]) // 报错但是可以正常运行的
}
