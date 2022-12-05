package main

import (
	"fmt"
)

/*1. 概念
现实生活中我们已经见到过很多接口的例子，比如 USB 接口，各个厂商按照一定的规范生产，最后的产品都是互相通用的。

golang中的接口是类似的一种抽象概念，它定义了对象的行为规范，只定义规范而不实现。
接口中定义的规范由具体的对象来实现。
通俗来讲，接口就是一个标准，他是对一个对象的行为和规范进行约定。如果一个对象要实现接口，则必须遵循接口的规范。
接口体现了程序设计的多态和高内聚低耦合的思想。
duck-typing: 当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子

接口是一组method的集合，接口中不能包含任何变量。

golang中的接口也是一种数据类型，不需要显式实现；只需要一个变量含有接口类型中的所有方法，那么这个变量就实现了这个接口。
*/
type animal interface {
	scream()
	move()
}

type bird struct {
	animal  // 结构体拥有了接口中的方法，需要具体实现才能调用，否则会报空指针错误
	voice   string
	feather string
}

func (b bird) scream() {
	fmt.Println(b.voice)
}

func (b bird) move() {
	fmt.Println("fly away!")
}

/******************************************************************/
type hunter struct{}

/*2.本质
接口本身也是一种数据类型，可以作为函数参数
*/
// hunt 需要一个animal类型的参数，凡是实现了该接口的对象(变量)，都可以作为参数传递过来
func (h hunter) hunt(a animal) {
	a.scream()
	a.move()
}

/******************************************************************/
/*3.空接口
go中的接口可以不定义任何方法，这种接口称为空接口,即包含0个method的interface
空接口表示任何约束，因此任何类型变量都可以实现空接口
*/

//空接口在实际项目中用的是非常多的，用空接口可以表示任意数据类型
var x interface{} = 123
var y interface{} = true

//空接口可以作为函数的参数，表示此函数可以接受任意类型的数据
func emptyInterfaceInfo(i interface{}) {
	fmt.Printf("value:%v ,type:%T \n", i, i)
}

//利用空接口，可以让map或者slice的值包含任意类型
var m map[string]interface{}

//var s []interface{}

/******************************************************************/
/*4.类型断言
一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。

如果我们想要判断空接口中值的类型，可以使用类型断言，语法格式为：x.(T)
其中，x表示类型为interface{}的变量；T表示断言x可能是的类型。
该语法返回两个参数，第一个参数是x转换为T类型后的变量，第二个值是一个布尔值，true表示断言成功，false表示断言失败。

类型断言实际项目中也是很常见的，比如要在一个统一的接口中处理各种不同类型的数据

类型断言除了作用于空接口，对于实现了一个非空接口的全部对象，也可以进行断言。
*/
func affirm(i interface{}) {
	s, ok := i.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("param is not a string")
	}
}

func handleType(i interface{}) {
	switch i.(type) { // 可以获取到interface的type,这个语法只能用在switch中
	case string:
		fmt.Println("param is a string")
	case int:
		fmt.Println("param is an int")
	}
}

/*******************************************************************/
/*5.结构体值接收者和指针接收者实现接口的区别
5.1 值接收者：
如果结构体中的方法是值接收者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量；
5.2 指针接收者：
如果结构体中的方法是指针接收者，那么只有实例化之后的结构体指针类型可以赋值给接口变量。
*/
// value receiver
type cat struct{}

func (c cat) scream() {
	fmt.Println("喵呜~")
}
func (c cat) move() {
	fmt.Println("skip high!")
}

// ptr receiver
type dog struct{}

func (d *dog) scream() {
	fmt.Println("汪汪汪！")
}

func (d *dog) move() {
	fmt.Println("run fast!")
}

/*******************************************************************/
/*
6.带有参数和返回值的接口
*/
type animalBase interface {
	setName(name string)
	getName() string
}
type fish struct {
	name string
}

// 这里的接收者必须是指针类型，因为结构体是值类型，避免修改成结构体的副本
func (f *fish) setName(name string) {
	f.name = name
}

func (f *fish) getName() string {
	return f.name
}

/*一个结构体也可以实现多个接口，这部分比较简单，自行测试即可*/
/*******************************************************************/
/*7.接口嵌套
接口间可以通过相互嵌套来实现新的接口
如果一个对象要实现这个新接口，那它就需要实现所有嵌套的接口
*/
type animalMgr interface {
	animalBase
}
type wale struct {
	name string
}

func (w *wale) setName(name string) {
	w.name = name
}

func (w *wale) getName() string {
	return w.name
}

/*******************************************************************/
/*8.空接口和类型断言的使用细节
当一个空接口传入切片或者结构体值时，不能直接获取切片内容或者结构体字段，
必须结合类型断言才可以。

11.15号 嘿今天写代码还真遇到这个问题
大概就是dcache中存放了一个字符串类型的切片 []string 解析之后要返回给客户端
*/
func assertionDetail(i interface{}) {
	// 先试试能不能直接断言成 []string
	fmt.Println(i.([]string)) //woc 竟然可以，牛b
	name := i.([]string)
	fmt.Println(name[1])
	var a2 []string
	a2 = append(a2, name...)
	fmt.Println(a2[2])
}

/*******************************************************************/
func main() {
	//sqarrow := bird{
	//	feather: "grey",
	//	voice:   "吱吱吱~",
	//}
	//sqarrow.scream()
	//sqarrow.move()

	//h := hunter{}
	//h.hunt(bird{})

	//空接口
	//fmt.Printf("value:%v ,type:%T \n", x, x)
	//fmt.Printf("value:%v ,type:%T \n", y, y)
	//s := "hello"
	//emptyInterfaceInfo(s)
	//
	//num := 123
	//emptyInterfaceInfo(num)

	// 类型断言
	//affirm(s)
	//affirm(num)
	//handleType(s)
	//handleType(num)

	// 结构体值接收者实现接口
	//prince := cat{}
	//h.hunt(prince)
	//h.hunt(&prince)

	// 结构体指针接收者实现接口
	//dog1 := &dog{}
	//h.hunt(dog1)

	// 带有参数和返回值的接口
	//f1 := &fish{
	//	name: "小卡",
	//}
	//var ab animalBase = f1 // 因为f1已经实现了接口，所以可以赋值给animalBase类型的ab
	//ab.setName("kaka")
	//fmt.Println(ab.getName())

	// 接口嵌套
	//w1 := &wale{
	//	name: "messi",
	//}
	//var am animalMgr = w1
	//am.setName("leo messi")
	//fmt.Println(am.getName())

	// 8.类型断言使用细节
	a1 := []string{"路飞", "索隆", "山治"}
	assertionDetail(a1)
}
