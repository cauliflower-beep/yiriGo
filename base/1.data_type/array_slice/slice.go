package a_s

import (
	"fmt"
)

/*
关键词：golang slice两个冒号的理解
*/

// func main() {
// 	// 所有空切片指向的地址都是一致的
// 	s1 := make([]int, 0)
// 	s2 := make([]int, 0)
// 	// 获取切片底层对应的数据结构
// 	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s1)), "\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
//
// 	// nil切片的地址是空的
// 	var s3 []int
// 	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&s3)))
// }

/*
切片是对数组的引用
*/

func ModifyArr() {
	arr := [5]int{1, 2, 3, 4, 5}

	s1 := arr[1:]
	fmt.Printf("地址关系|arr:%p|s1:%p|s1[0]:%p\n", &arr, &s1, &s1[0])
	s2 := s1[:2]
	fmt.Printf("[before modify]|arr:%+v|s1:%+v|s2:%+v\n", arr, s1, s2)

	s3 := arr[:]
	fmt.Println("s3:", s3)

	s2[0] = 6
	fmt.Printf("[after modify]|arr:%+v|s1:%+v|s2:%+v\n", arr, s1, s2)

	fmt.Println("s3:", s3)
}

func expand() {
	arr := [5]int{1, 2, 3}
	sli := arr[2:5]
	fmt.Printf("arr:%p|sli:%p\n", &arr, &sli)
	fmt.Printf("[before expand sli info]|length:%d|cap:%d\n", len(sli), cap(sli))
	// 查看底层数组的指针变化
	fmt.Printf("扩容前 sli 底层数组地址:%p\n", sli)

	sli = append(sli, 4)
	fmt.Printf("[when expand sli info1]|length:%d|cap:%d\n", len(sli), cap(sli))
	fmt.Printf("扩容中 sli 底层数组地址1:%p\n", sli)

	sli = append(sli, 5)
	fmt.Printf("arr:%p|sli:%p\n", &arr, &sli)
	fmt.Printf("[when expand sli info2]|length:%d|cap:%d\n", len(sli), cap(sli))
	fmt.Printf("扩容中 sli 底层数组地址2:%p\n", sli)

	sli = append(sli, 6)
	sli = append(sli, 7)
	sli = append(sli, 8)
	sli = append(sli, 9)
	sli = append(sli, 10)
	// 扩容超出最大容量之后，切片本身地址并没有发生变化，它只是一个拥有cap、len、指向底层数组指针的一个结构体
	fmt.Printf("arr:%p|sli:%p\n", &arr, &sli)
	fmt.Printf("[after expand sli info]|length:%d|cap:%d\n", len(sli), cap(sli))
	// 指向底层数组的指针会发生变化，因为会为它创建一个新的数组
	fmt.Printf("扩容后 sli 底层数组地址:%p\n", sli)

}
