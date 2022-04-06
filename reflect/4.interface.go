package main

import (
	"fmt"
	"reflect"
)

/*
反射可以在运行时动态获取程序的各种详细信息
反射获取interface类型信息
 */

/*
类型与种类的区别
开发时，使用最多的是类型type,但在反射中，当需要区分一个大品种的类型时，就会用到种类kind.
例如，需要统一判断类型中的指针时，使用类king信息就较为方便.
 */
func reflect_Type(a interface{}){
	t := reflect.TypeOf(a)		// 类型
	fmt.Println("type is",t)

	k := t.Kind()		// 种类
	switch k{
	case reflect.Float64:
		fmt.Println("kind is float64")
	case reflect.String:
		fmt.Println("kind is string")
	default:
		fmt.Println("kind is",k )
	}
}
