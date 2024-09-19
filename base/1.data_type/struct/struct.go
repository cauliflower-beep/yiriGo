package _struct

import (
	"fmt"
)

/********************************定义********************************/

type Player struct {
	id            int
	nickname, sex string
	age           int
	pet           Bird
}

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

/***********************值传递***************************/

func getName(player Player) string {
	player.nickname = "野原美伢"
	return player.nickname
}

func getAge(player *Player) int {
	player.nickname = "野原广志"
	return player.age
}

/**************************Tag***********************/

type Bird struct {
	Id     int    `json:"id" curd:"AUTO_INCREMENT"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Weight int    `json:"weight"`
	Gender string `json:"gender"`
}

/******************************method****************************/

func (p *Player) setName(name string) {
	fmt.Println("newname:", name)
	p.nickname = name
}

/***************************main***************************/
func s() {
	// 2.assign
	var p1 = Player{
		nickname: "小花椰种植员",
		age:      16,
	}
	var fly = Bird{
		Id:    0,
		Name:  "fly",
		Color: "blue",
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
