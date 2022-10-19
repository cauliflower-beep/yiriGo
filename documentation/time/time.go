package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 24小时之内的计算
	t1, _ := time.ParseDuration("12h")
	m1 := now.Add(t1).Unix()
	m2 := uint32(m1)
	fmt.Println(m1, m2)
}
