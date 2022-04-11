package main

import (
	"context"
	"fmt"
	"time"
)
/*
日常开发时，为了完成一个复杂的需求，可能会开多个goroutine去做一些事情
为了更方便的控制这些在一个请求中开启的多个goroutine，可以使用withCancel来衍生一个context传递到不同的goroutine中
当我们想让这些goroutine停止运行，就可以调用cancel来进行取消
 */
func main(){
	ctx,cancel := context.WithCancel(context.Background())
	go speak(ctx)
	time.Sleep(10*time.Second)
	cancel()
	time.Sleep(time.Second)
}

func speak(ctx context.Context){
	for range time.Tick(time.Second){
		select {
		case <-ctx.Done():
			fmt.Println("shut your mouth!")
			return
		default:
			fmt.Println("lalala~")
		}
	}
}