package main

import (
	"fmt"
)

/*
go中的 struct 与数组一样，属于复合类型，并非引用类型
与其他面向对象编程语言中的 class 类似，struct可以定义字段（属性）和方法
*/

/////////////////////////////////////////////////////////////////
// 1.1define
/*
同类型的字段可以放在一行定义，但比较好的编程习惯是分行定义
struct可以通过复合其他结构体来构建更复杂的结构体，但不能包含自己
*/
type Player struct {
	id            int
	nickname, sex string
	age           int
	pet           Bird
}

/* 1.2 匿名字段
上面定义的 Player结构体，字段名与类型是一一对应的。
实际上Go也支持只提供类型，而不写字段名的方式，称之为匿名字段，也叫嵌入字段。
*/

// scene 自定义类型
type scene string
type Game struct {
	// 匿名字段为struct
	Player // Game 包含Player结构体的所有字段，可通过 Game.id 或 Game.Player.id 的方式访问
	level  int
	int        // 匿名字段为内置类型 Game.int
	scene      // 匿名字段为自定义类型
	id     int // Game 拥有与Player结构体同名的字段，Game.id会优先访问Game的字段，而不是Player里面的字段
}

/////////////////////////////////////////////////////////
/*
结构体与数组一样，都是值传递.
比如当把数组或结构体作为实参传给函数的形参时，会复制一个副本
所以为了提高性能，一般不会把数组直接传递给函数，而是使用切片(引用类型)代替，而把结构体传给函数时，可以使用指针结构体

！！数据量小的结构体，传拷贝比传指针快，但如果结构体里面有map，传指针一定比传拷贝要快
*/

func getName(player Player) string {
	player.nickname = "野原美伢"
	return player.nickname
}

func getAge(player *Player) int {
	player.nickname = "野原广志"
	return player.age
}

///////////////////////////////////////////////////////////////////////////////

/*
Tags
在定义结构体字段时，除字段名称和数据类型外，还可以使用反引号为结构体字段声明元信息
这种元信息称为Tag，用于编译阶段关联到字段当中

Tag由反引号括起来的一系列用空格分隔的 key:"value" 组成.
*/
type Bird struct {
	id     int    `json:"id" curd:"AUTO_INCREMENT"`
	name   string `json:"name"`
	color  string `json:"color"`
	weight int    `json:"weight"`
	gender string `json:"gender"`
}

//////////////////////////////////////////////////////////////////////////////

/*
方法：
在Go语言中，将函数绑定到具体的类型中，则称该函数是该类型的方法，其定义的方式是在func与函数名称之间加上具体类型变量，这个类型变量称为方法接收器
并不是只有结构体才能绑定方法，任何类型都可以绑定方法，只是我们这里介绍将方法绑定到结构体中。
*/

/*
因为是值传递，所以我们指定结构体为方法接收器时，通常传入结构体指针，
否则函数操作的只是结构体的一个副本，并不会对原结构体造成影响
*/
func (p *Player) setName(name string) {
	fmt.Println("newname.", name)
	p.nickname = name
}

/////////////////////////////////////////////////////////////////////////////////
func main() {
	// 2.assign
	var p1 = Player{
		nickname: "小花椰种植员",
		age:      16,
	}
	var fly = Bird{
		id:    0,
		name:  "fly",
		color: "blue",
	}
	var p2 = Player{1, "广志", "male", 35, fly}

	// 3.访问字段
	fmt.Println("访问普通字段：", p1.nickname, p2.age)
	//fmt.Println(p2.age)
	// 3.1 访问匿名字段
	round1 := Game{Player{1, "广志", "male", 35, fly}, 1, 1, "都市", 0}
	fmt.Println("匿名字段访问：", round1.nickname, round1.Player.nickname, round1.id)

	// 4.值传递
	fmt.Println("值传递展示：", getName(p1), p1) // 由于是值传递，函数内部对结构体字段的修改，并不会影响原有的数据

	// 5.结构体指针
	fmt.Println("指针传递结果展示：", getAge(&p2), p2) // 指针传递会影响原来的值

	// 6.方法
	p1.setName("风间彻")
	fmt.Println("方法调用展示：", p1)

}
