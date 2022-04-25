package main

import (
	"fmt"
	"github.com/wxnacy/wgo/arrays"
	"time"
)

/*
判断slice中是否含有某个元素
 */

//var second =[]int{8,12,16,20,24,28,30,32,36,40,44,48,52,56,0}
var second = []int{}
func assignSecond(){
	for i := 0;i<60;i++{
		second = append(second,i)
	}
}
var hour = 19
func main(){
	assignSecond()
	ticker := time.NewTicker(1*time.Second)
	for {
		<- ticker.C
		if arrays.Contains(second,time.Now().Second()) != -1{
			fmt.Printf("当前是第%d秒\n",time.Now().Second())
		}
	}
	ticker.Stop()
}