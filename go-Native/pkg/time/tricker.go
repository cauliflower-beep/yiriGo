package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	tricker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-tricker.C:
				fmt.Println(t.Format("2006-1-2 15:04:05"))
			}
		}
	}()

	select {}
}
