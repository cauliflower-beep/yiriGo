package main

import "fmt"

func main() {
	test()
}

func test() {
	str := "Hello"
	// Printf format %d has arg str of wrong type string
	fmt.Printf("%d World", str)
}
