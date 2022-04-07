package main

import (
	"fmt"
)

/*
注意！日常开发时，还是要注意尽量避免使用指针类型的通道！
 */

/*
golang中使用for range语句进行迭代非常的便捷，
但在涉及到指针时就得小心一点了
 */

func iterate1(num []int,ch chan *int){
	for _,v := range num{
		ch <- &v
		//time.Sleep(time.Second)
	}
	close(ch)
}

func iterate2(num []int,ch chan *int){
	for i,_ := range num{
		// 1.引入中间变量
		//temp := v
		//ch <- &temp

		// 2. 直接引用数据内存
		ch <- &num[i]
	}
	close(ch)
}

func main(){
	ch := make(chan *int,5)

	//sender
	num := []int{1,2,3,4,5}

	//go iterate1(num,ch)

	go iterate2(num,ch)
	// receiver
	for v:= range ch {
		fmt.Println(*v,v)
	}
}
/*
iterate1中，结果出来会是 5 5 5 5 5.
原因：
可以将打印内容改为 fmt.Println(v),发现打印的地址均为 0xc000012090
在for range语句中，v变量用于保存迭代 num 数组所得的值，但是v只被声明了一次,
此后都是将迭代 num 出的值赋值给v,v 变量的内存地址始终未变，这样再将 v 的地址发送给 ch 通道，发送的都是同一个地址，所以打印出来的结果都是相同的

若想打印 1 2 3 4 5
可以引入一个中间变量，每次迭代都重新声明一个变量 temp , 赋值后再将其地址发送给 ch;
或者直接引用数据的内存（推荐，无需开辟新内存）
参考iterate2.
 */
