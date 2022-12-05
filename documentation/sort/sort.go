package main

import (
	"fmt"
	"sort"
)

/*
	使用 golang 的 sort 包来排序，
*/

func main() {
	// 数字数组排序
	ints := []int{11, 44, 33, 22}
	sort.Ints(ints)                              //默认升序
	fmt.Printf("%v\n", ints)                     //[11 22 33 44]
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) //降序排序
	fmt.Printf("%v\n", ints)

	// 字符串数组排序
	str := []string{"apple", "lemen", "banana", "fruit"}
	sort.Strings(str)
	fmt.Printf("%v\n", str)                        //默认升序
	sort.Sort(sort.Reverse(sort.StringSlice(str))) //降序排序
	fmt.Printf("%v\n", str)

	// 切片排序
	slices := []int{11, 11, 44, 55, 11, 44}
	sort.Slice(slices, func(i, j int) bool {
		//return slices[i] < slices[j] //升序  即前面的值比后面的小
		return slices[i] > slices[j] //降序  即前面的值比后面的大
	})
	fmt.Printf("%v\n", slices) //[55 44 44 11 11 11]

	// 结构体自定义排序
	type student struct {
		name string
		age  int
	}

	s := []student{{"babala", 52}, {"anly", 50}, {"babala", 51}}
	sort.Slice(s, func(i, j int) bool {
		if s[i].name == s[j].name { //如果名字相同 按照年龄熊大到小
			return s[i].age > s[j].age // 年龄降序
		}
		return s[i].name < s[j].name // 名字升序
	})
	fmt.Printf("%v\n", s) //[{anly 50} {babala 52} {babala 51}]
}
