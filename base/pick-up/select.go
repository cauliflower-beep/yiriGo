package main

import (
	"fmt"
	"sync"
)

/*
使用方法
select 语句用于监控并选择一组case语句执行相应的代码。
他看起来类似于 switch 语句，但是所有 case 中的表达式都必须是 channel 的发送或接收操作。
*/
var selectWg sync.WaitGroup

func selectDemo(ch1, ch2 chan int) {
	select {
	case <-ch1:
		fmt.Println("ch1 ...")
		break
	case <-ch2:
		fmt.Println("ch2 ...")
		break
	default:
		fmt.Println("default process...")
	}
	selectWg.Done()
}

/*
	死锁问题
	如果你执行过 ./block.go 会发现main中的 _goroutine 执行完之后，终端报了死锁的错误，原因就是main中存在 select{}.
	select 为什么会存在死锁问题呢?答案都在spec中:https://go.dev/ref/spec#Select_statements
	翻译第一条：在输入"select"语句时，接收操作的通道操作数以及发送语句的通道和右侧表达式仅按源顺序计算一次。
	意味着 select 会拿到右侧语句的结果才执行，如果拿不到就会一直堵塞，造成死锁。
*/
/*
	可能上述的解释还是不太清楚，来看个例子:
	demo1: 在进入select语句时，会按源码的顺序对每一个 case 子句进行求值：这个求值只针对发送或接收操作的额外表达式。
	无论 select 最终选择了哪个 case，getVal() 都会按照源码顺序执行： getVal(1) 和 getVal(2)，也就是它们必然先输出：
	getVal, i= 1 getVal, i= 2
*/
func demo1() {
	ch := make(chan int)
	go func() {
		select {
		case ch <- getVal(1):
			fmt.Println("in first case")
		case ch <- getVal(2):
			fmt.Println("in second case")
		default:
			fmt.Println("default process...")
		}
	}()

	fmt.Println("The val:", <-ch)
}

func getVal(i int) int {
	//_panic("im done")
	fmt.Println("getVal, i=", i)
	return i
}

func main() {

	// 使用方法
	selectWg.Add(1)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go selectDemo(ch1, ch2)
	//ch2 <- 1
	selectWg.Wait()

	// 死锁问题
	//demo1()
}
