package main

import "fmt"

func main() {
	/*
		可以使用如下的方法给 fmt.Println 输出着色
	*/
	//colorReset := "\033[0m"
	//colorRed := "\033[31m"
	//fmt.Println(string(colorRed), "圣诞节快乐！", string(colorReset))
	//fmt.Println("吃苹果了嘛？！")
	fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "圣诞节快乐！", 0x1B)
}
