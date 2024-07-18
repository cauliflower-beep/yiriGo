package main

import (
	"fmt"
	"reflect"
)

var sn1 = struct {
	age  int
	name string
}{age: 16, name: "conan"}

var sn2 = struct {
	name string
	age  int
}{age: 16, name: "conan"}

var sn3 = struct {
	age  int
	name string
}{age: 16, name: "conan"}

var sm1 = struct {
	name  string
	grade map[string]int
}{name: "tom", grade: map[string]int{"math": 92, "english": 93}}

var sm2 = struct {
	name  string
	grade map[string]int
}{name: "tom", grade: map[string]int{"math": 92, "english": 93}}

func main() {
	//fmt.Println(sn1 == sn2) // 结构体字段顺序不同 不能比较
	if reflect.DeepEqual(sn1, sn2) {
		fmt.Println("sn1 == sn2")
	} else {
		fmt.Println("sn1 != sn2")
	}

	fmt.Println(sn1 == sn3) // true

	//fmt.Println(sm1 == sm2) // 含有map字段 不能直接比较
	// reflect.DeepEqual 比较
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	} else {
		fmt.Println("sm1 != sm2")
	}
}
