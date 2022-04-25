package main

import (
	"fmt"
	"time"
)

func getNumber() int64{
	return int64(32)
}

func main(){
	now := time.Now().Hour()
	fmt.Println("现在",now,"点了")
	switch {
	case now <= 8:
		fmt.Println("距离今天第一轮开奖还有：",8-now,"h")
	case now <= 12:
		fmt.Println("距离今天第二轮开奖还有：",12-now,"h")
	case now <= 16:
		fmt.Println("距离今天第三轮开奖还有：",16-now,"h")
	case now <= 20:
		fmt.Println("距离今天第四轮开奖还有：",20-now,"h")
	case now <= 24:
		fmt.Println("距离今天第五轮开奖还有：",24-now,"h")
	}
	t := getNumber()
	d := time.Duration(t)
	fmt.Printf("t = %d,type is %t\n",t,t)
	fmt.Printf("d = %d,type is %t\n",d,d)
}
