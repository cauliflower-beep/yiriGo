package _struct

import "testing"

// go test -run TestEmbedInterface
func TestEmbedInterface(t *testing.T) {
	fcb2016 := FCB{
		Players:    []string{"梅西", "苏亚雷斯", "内马尔"},
		Instructor: "巴尔韦德",
	}

	famousClub(fcb2016) // FCB 嵌套了 footballClub 接口 所以可以直接作为该接口类型传入方法中

	fcb2016.GetInstructor() // 可以重写接口方法

	// 实现了接口的对象也可以注入
	fcbw := FCBWomen{}
	fcb2017 := FCB{
		footballClub: fcbw,
		Players:      []string{"梅西", "苏亚雷斯", "内马尔"},
		Instructor:   "111",
	}

	fcb2017.GetPlayers()
}
