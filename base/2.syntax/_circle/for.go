package _circle

import "fmt"

func general() {
	arr := [...]int{999999: 1}
	fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
	}
	//fmt.Println("general done.")
}

func otherWay() {
	arr := [...]int{999999: 1}
	fmt.Println(len(arr))
	for i, al := 0, len(arr); i < al; i++ {
	}
	//fmt.Println("otherWay done.")
}
