package main

import "sync"

var (
	lock sync.Mutex
	wg 	 sync.WaitGroup
)

var num int64 = 0
