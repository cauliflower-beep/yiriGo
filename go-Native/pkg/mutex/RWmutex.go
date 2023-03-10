package main

import (
	"fmt"
	"sync"
)

var (
	rwLock sync.RWMutex
)

func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println(num)
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	num++
	rwLock.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}

	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
}
