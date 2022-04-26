package main

import "fmt"

/*
1.new()
new的函数声明:	func new(Type)*Type
new 是Go的内建函数，用于分配内存。其中，第一个参数是类型，返回值是类型的指针，其值被初始化为 "零值" .不同类型的零值是不同的.
来看一个例子：
*/

func newDemo() {
	id := new(int)
	name := new(string)
	gotMarried := new(bool)

	fmt.Printf("id type:%T,id value:%d\n", id, *id)
	fmt.Printf("name type:%T,name value:%v\n", name, *name)
	fmt.Printf("gotMarried type:%T,gotMarried value:%v\n", gotMarried, *gotMarried)
}

/*
2.make()
make的函数声明:	func make(t Type,size ...IntegerType) Type
make 也是Go的内建函数，仅用于分配和初始化slice/map/channel 类型的对象，三种类型都是结构。返回值为类型，而不是指针。
 */

/*
3.new和make的区别：
new和make都用于分配内存；
new和make都是在堆上分配内存；
new对指针类型分配内存，返回值是分配类型的指针，new不能直接对slice、map、channel分配内存；
make仅用于slice、map、channel的初始化，返回值为类型本身，而不是指针
 */
func main(){
	newDemo()
}