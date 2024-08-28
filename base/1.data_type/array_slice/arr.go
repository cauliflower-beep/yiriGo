package a_s

import "fmt"

/*
1.数组定义的几种方式:
*/
func define() {
	var names [3]string
	fmt.Println("1.默认初始化方式|", names)

	/*
		数组长度在定义之后无法再次修改 ；
		{}中的元素个数不能大于[]中定义的数组长度，超出长度后继续追加元素会报错
	*/
	var balance = [5]int{1, 2, 3}
	fmt.Println("2.初始化数组并赋值|", balance)

	// 若数组长度位置出现"..."，表示数组的长度是根据初始化值的个数计算的
	q := [...]int{1, 2, 3}
	fmt.Println("3.函数内初始化|", q)

	// 下面定义一个数组，长度为10，前9项都是0，最后一项是1
	r := [...]int{9: 1}
	m := [...]int{0: 1, 2: 2, 5: 3}
	fmt.Printf("4.特殊方式初始化|r:%v^length:%d|%v\n", r, len(r), m)

	x, y := 1, 2
	a := [...]*int{&x, &y}
	fmt.Println("5.指针数组初始化|", a)
}

/*
2.数组是值类型，作为函数参数传递时，会copy出一个副本
对副本的修改不会影响原有数组
*/
func setDataCopy(arr [5]int) {
	arr[1] = 1
}
func setDataAddr(arr *[5]int) {
	/*
		arr 是一个指向[5]int类型数组的指针
		直接使用arr[1]访问数组元素的时候，Go编译器会自动解引用指针来访问它指向的数组
		这是一种语法糖，可以是代码更简洁易读
	*/
	arr[1] = 1
	//(*arr)[1] = 2 // 显示解引用
}

/*
3.数组指针地址与数据第一个值地址相同
*/
func arrAddr() {
	arr := [3]int{1, 2, 3}
	fmt.Printf("arrAddr:%p|arr[0]Addr:%p|arr[1]Addr:%p\n", &arr, &arr[0], &arr[1])
}
