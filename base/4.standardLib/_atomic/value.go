package _atomic

import (
	"fmt"
	"sync/atomic"
)

var v atomic.Value

func store() {
	c := make(chan struct{})
	fmt.Println(c)
	v.Store(c)

	c0 := v.Load()
	fmt.Println(c0)

	c1 := v.Load()
	fmt.Println(c1)

	c2 := v.Load()
	fmt.Println(c2)

}
