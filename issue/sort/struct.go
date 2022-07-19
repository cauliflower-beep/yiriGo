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
	sort.Sort(persons(arr))
	fmt.Println(arr)
}
