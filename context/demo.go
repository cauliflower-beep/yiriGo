package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
有一份在线意见征集表，限时时间内可以由多个goroutine共享读写
时间到了之后需要立即停止并退出访问它的goroutine
 */

var file = "is the beauty kill the beast"

var rwMutex sync.RWMutex

func read(ctx context.Context){
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	rand.Seed(time.Now().Unix())
	id := rand.Intn(1000)
	for range time.Tick(time.Second){
		select {
		case <-ctx.Done():
			fmt.Println("Time is up, sharing ends,good bye ",id)
			return
		default:
			fmt.Println(id,"is reading the file")
		}
	}
}

func main(){
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	for i :=0;i<2;i++{
		go read(ctx)
		time.Sleep(time.Second)
	}
	select {}	// 阻塞等待
}