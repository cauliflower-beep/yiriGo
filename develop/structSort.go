package main

import (
	"fmt"
	"sort"
)

/*
	业务场景：
	从跨表分页中取出了若干记录，映射到了一个结构体数组中。
	现在需要按照其中的time字段升序排序，取出最早的前10条记录
*/

// 问题：如何依据结构体中的某个字段进行排序？

// 按年纪从大到小排序
type person struct {
	age    int
	height int
}

type persons []person

func (p persons) Len() int { return len(p) }

func (p persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p persons) Less(i, j int) bool { return p[i].age > p[j].age }

func main() {
	arr := []person{
		person{10, 12},
		person{11, 12},
		person{9, 12},
		person{14, 12},
		person{7, 12},
		person{7, 12},
		person{22, 12},
	}
	// 22.07.14
	sort.Sort(persons(arr))
	fmt.Println("sort.Sort  res|", arr)

	// 22.08.12
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].age < arr[j].age
	})
	fmt.Println("sort.Slice res|", arr)
}

// 一组接口，实现多种结构体排序	8.05
//type obj interface{}
//
//type objs []obj
//
//func (ob objs) Len() int { return len(ob) }
//
//func (ob objs) Swap(i, j int) { ob[i], ob[j] = ob[j], ob[i] }
//
//func (ob objs) Less(i, j int) bool {
//	//if ob.(persons) == persons{
//	//
//	//}
//	//return ob.(persons)[i].age > ob.(persons)[j].age
//	return false
//}
