package main

import (
	"fmt"
	"reflect"
)

/*
反射可以在运行时动态获取程序的各种详细信息
反射获取interface类型信息
*/

//reflect_Type 参数为一个空接口，意味着可以接受任意类型的值
func reflect_Type(a interface{}) {
	t := reflect.TypeOf(a) // 类型
	fmt.Println(a, "type is", t)

	/*
		对于返回值 t 还有两个重要方法：kind() 和 Name()
	*/

	k := t.Kind() // 种类
	/*
		类型与种类的区别
		开发时，使用最多的是类型type,但在反射中，当需要区分一个大品种的类型时，就会用到种类kind.
		例如，需要统一判断类型中的指针时，使用类kind信息就较为方便.
	*/
	switch k {
	case reflect.Float64:
		fmt.Println("kind is float64")
	case reflect.String:
		fmt.Println("kind is string")
	default:
		fmt.Println("kind is", k)
	}

	n := t.Name() //类型名称
	fmt.Println(a, "name is", n)
}

/*
测试反射是否可以获取自定义类型以及结构体类型
*/
type myInt int
type Person struct {
	name string
	age  int
}

func main() {
	// get type of interface
	a := 3.64
	reflect_Type(a)
	b := "hello reflect!"
	reflect_Type(b)
	c := make(map[string]string)
	reflect_Type(c)
	var d myInt = 3
	reflect_Type(d)
	e := Person{
		name: "zxx",
		age:  21,
	}
	reflect_Type(e)

}
