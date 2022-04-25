package main

import (
	"fmt"
	"time"
)

// now := time.now(),时间变动，now值会变吗？测试发现，不会变的

func main(){

	tricker := time.NewTicker(1*time.Second)
	for {
		<-tricker.C
		now := time.Now()
		fmt.Println(now.Second(),now.Hour())
	}
}
