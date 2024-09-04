package _func

import (
	"fmt"
	"testing"
)

func TestDefine(t *testing.T) {
	prince := film{
		title: "小王子",
		date:  "2014-06-28",
	}
	prince.release()
}

// 值接收者
func TestValReceiver(t *testing.T) {
	// 值接收者 值调用
	redPhoenix := bird{
		color: "red",
	}
	fmt.Println("raw color", redPhoenix.color)
	redPhoenix.changeColor("yellow")
	fmt.Println("new color", redPhoenix.color)

	// 值接收者 指针调用
	yellowPhoenix := &bird{
		color: "yellow",
	}
	fmt.Println("raw color", yellowPhoenix.color)
	yellowPhoenix.changeColor("red")
	fmt.Println("new color", yellowPhoenix.color)
}

// 指针接收者
func TestPtrReceiver(t *testing.T) {
	// 值调用
	fm1 := film{
		title: "red",
	}
	fmt.Println("raw title", fm1.title)
	fm1.changeTitle("yellow")
	fmt.Println("new title", fm1.title)

	// 值接收者 指针调用
	fm2 := &film{
		title: "yellow",
	}
	fmt.Println("raw title", fm2.title)
	fm2.changeTitle("red")
	fmt.Println("new title", fm2.title)
}
