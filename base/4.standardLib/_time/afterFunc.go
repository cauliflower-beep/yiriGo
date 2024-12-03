package main

import (
	"fmt"
	"time"
)

func main() {
	_ = time.AfterFunc(time.Second*5, func() {
		fmt.Println("over")
	})
	time.Sleep(10 * time.Second)
}
