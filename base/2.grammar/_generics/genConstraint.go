package main

import (
	"fmt"
	"strconv"
)

/***********************1. 使用 interface 中规定的类型约束泛型函数的参数*********************/

type NumStr interface {
	// "|" 表示取并集 如果传入的参数不在集合限制范围内，就会报错
	Num | Str
}
type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~complex64 | ~complex128
}
type Str interface {
	string
}

func add[T NumStr](a, b T) T {
	return a + b
}

/***********************2. 使用 interface 中规定的方法约束泛型函数的参数*********************/

type Price int

func (p1 Price) String() string {
	return strconv.Itoa(int(p1))
}

type Price2 string

func (p2 Price2) String() string {
	return string(p2)
}

type ShowPrice interface {
	String() string
	~int | ~string
}

func ShowPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return
}

/***********************2. 使用 interface 中规定的方法和类型来双重约束泛型函数的参数*********************/
type Price3 int
type ShowPrice3 interface {
	String() string
	int | string
}

func (p3 Price3) String() string {
	return strconv.Itoa(int(p3))
}
func ShowPriceList2[T ShowPrice3](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return
}

/****************************main**************************/
func main() {
	fmt.Println(add(3, 4))
	fmt.Println(add("hello", "world"))

	fmt.Println(ShowPriceList([]Price{1, 2}))
	fmt.Println(ShowPriceList([]Price2{"a", "b"}))

	fmt.Println(ShowPriceList([]Price3{1, 2}))
}
