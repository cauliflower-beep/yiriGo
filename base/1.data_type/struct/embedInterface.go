package _struct

import "fmt"

/*
结构体嵌套接口，那么这个结构体本身就是这个接口的一个实现，
即使该结构体本身并没有实现接口中的任何一个方法
标准库中的context.cancelCtx就是这种形式的一个典型应用
*/

type footballClub interface {
	GetPlayers()    // 获取球员
	GetInstructor() // 获取教练
}

type FCB struct {
	footballClub
	Players    []string
	Instructor string
}

func famousClub(fc footballClub) {
	fc.GetInstructor()
}

func (f FCB) GetInstructor() {
	fmt.Println(f.Instructor)
}

/*
实现了接口的对象，也可以注入
*/

type FCBWomen struct{}

func (fcb FCBWomen) GetPlayers() {}

func (fcb FCBWomen) GetInstructor() {}
