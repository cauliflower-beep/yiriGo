package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 十进制
	var a = 10
	a2 := strconv.FormatInt(int64(a), 2) //10进制转换为2进制
	fmt.Printf("%v \n", a2)              //1010

	a8 := strconv.FormatInt(int64(a), 8) //10进制转换为8进制
	fmt.Printf("%v \n", a8)              //12

	a16 := strconv.FormatInt(int64(a), 16) //10进制转换为16进制
	fmt.Printf("%v \n", a16)               //a

	var b = "1010"
	p, _ := strconv.ParseInt(b, 2, 64) // 2进制转10进制
	fmt.Printf("%v\n", p)              //10

	var b2 = "345"
	p2, _ := strconv.ParseInt(b2, 8, 64) // 8进制转10进制
	fmt.Printf("%v\n", p2)               //229

	// 字符串<--->int
	i1, _ := strconv.Atoi("321")
	fmt.Println(i1 + 2)
	i2, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Println(i2, "converted failed!")
	}

	// 数字转字符串
	fmt.Println(strconv.FormatInt(1666936521, 10))

}
