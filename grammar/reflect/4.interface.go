package main

import (
	"fmt"
	"reflect"
)

/*
反射可以在运行时动态获取程序的各种详细信息
反射获取interface类型信息
*/

/*************************reflect.TypeOf***************************************************/
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

/****************************reflect.ValueOf****************************************/
//reflect_Value 参数为一个空接口，意味着可以接受任意类型的值
func reflect_Value(a interface{}) {
	/*
		var num = 10+a	//mismatched types untyped int and interface{}
		假如我们需要实现10+a这样的操作该怎么办呢？
	*/
	/*
		1.可以使用类型断言：
		var num = 10 + a.(int)
		fmt.Println(num)
	*/
	/*
		2.采用反射的方式
		v := reflect.ValueOf(a)
		// 反射获取变量的原始值
		n := v.Int() + 10
		fmt.Println(n)
	*/

	/*
		不同类型的原始值是怎样的呢？
		reflect.Value类型也有一个Kind()方法,可以通过它获取变量的种类信息
	*/
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		fmt.Println(a, "kind is float64")
	case reflect.String:
		fmt.Println(a, "kind is string")
	default:
		fmt.Println(a, "kind is", k)
	}

}

/*
demo 利用反射设置某个变量的值
*/
func reflectSetValue(x interface{}) {
	// *x = 120 // invalid indirect of x (type interface{})

	// v,_ := x.(*int)
	// *v = 120 // invalid memory address of nil pointer dereferenc

	v := reflect.ValueOf(x)

	fmt.Println(v.Elem().Kind()) //如果参数传递的是地址，就可以通过.Elem()获取类型
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(120)
	}
}

/*********************************main**********************************/
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

	/****value****/
	reflect_Value(3)
	reflect_Value("hello")
	reflect_Value(e)

	//set value
	var f int64 = 60
	reflectSetValue(&f)
	fmt.Println("f=", f)
}
