package main

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"
)

var wg4 sync.WaitGroup
var c = make(chan int)
var d = make(chan int)

func cat() {
	for i := 0; i < 10; i++ {
		fmt.Println("cat")
	}
	c <- 1
	wg4.Done()
}
func dog() {
	select {
	case <-c:
		for i := 0; i < 10; i++ {
			fmt.Println("dog")
		}
		d <- 1
		close(c)
		wg4.Done()
	}
}
func fish() {
	select {
	case <-d:
		for i := 0; i < 10; i++ {
			fmt.Println("fish")
		}
		close(d)
		wg4.Done()
	}
}

func context_test() {

	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		println(c.Value("a").(string))
	}(ctx)

	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				println("child process interrupt...")
				return
			default:
				println("enter default")
			}
		}
	}(timeoutCtx)

	time.Sleep(time.Second * 1)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(time.Second * 1)
		println("main process exit ")
	}

}
func main() {
	//x := 3.4
	//fmt.Println(int32(x))
	//wg4.Add(3)
	//go fish()
	//go dog()
	//go cat()
	//wg4.Wait()

	//s := "泛海控股发现股价创新低风险,,神州数码发现主力资金流出风险,"
	//s = strings.Trim(s, ",")
	//fmt.Println(s)
	//context_test()
	//time.Sleep(time.Second * 3)

	var b1 [5]byte
	var b2 []byte
	b3 := make([]byte, 1024)
	t_b1 := reflect.TypeOf(b1)
	t_b2 := reflect.TypeOf(b2)
	t_b3 := reflect.TypeOf(b3)
	fmt.Println(t_b1.Kind(), t_b2.Kind(), t_b3.Kind())
}
