package main

/*
说明：
iota是一个预先声明的标识符，用于在(通常用括号括起来)的const声明中表示当前const规范的非类型的化整数序数。iota是从0开始索引的。

从说明中可以知道：
	1.iota需要配合const才能使用，通常是用括号括起来的const
	2.iota是用来表示非类型的化整数序数，后续的常量会自增
	3.iota通常是从0开始索引的(需要注意iota在const中的位置，位置会影响iota的初始索引值，后面会说明)
*/

// 用法1：枚举
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

/* 用法2：常量表达式
iota可以参与 加、减、乘、除、位移 等表达式。
*/
type Flags uint

const (
	FlagOne = 1 << iota
	FlagTwo
	FlagThree
)

// 参与更复杂的表达式，例如：
const (
	_  = 1 << (10 * iota)
	KB // 1 << (10*1)
	MB // 1 << (10*2)
	GB // 1 << (10*3)
	TB // 1 << (10*4)
	PB // 1 << (10*5)
	EB // 1 << (10*6)
	ZB // 1 << (10*7)
	YB // 1 << (10*8)
)

/*
特别说明：
以上iota的使用通常位于const内部（带括号），且位于内部第一个变量声明处。如果不满足这些条件呢？

如下，出现多个iota的情况：
*/
const x = iota // 0
const y = iota // 0
const (
	z  = iota // 0
	z1        // 1
	_
	z3 = iota // 3
)

/*
从上，可以看出：

iota位于const头部时，从0开始
iota位于const非头部时，会受到iota的位置影响，iota初始值为位置顺序索引值
如下：
*/
const (
	a int     = 0    // 0
	b float64 = 0    // 0
	c         = iota // 2
	d         = "test"
	e         = iota // 4
	_
	f = iota + 1 // 7
)
