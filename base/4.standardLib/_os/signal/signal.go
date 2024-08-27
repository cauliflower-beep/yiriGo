package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	go func() {
		/*
			main 中可以启动一个goroutine，处理具体逻辑，减少main的臃肿
		*/
		goSignalChan := make(chan os.Signal, 1)
		signal.Notify(goSignalChan, os.Interrupt)
		<-goSignalChan
		fmt.Println("_goroutine close.")
	}()
	// 外部再创建一个信号监听器，接收到终止信号后打出日志方便调试
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("main close.")
}
