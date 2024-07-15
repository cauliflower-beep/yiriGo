package main

import "fmt"

func getFunc(flag string) func(x, y int) int {
	if flag == "add" {
		return func(x, y int) int {
			return x + y
		}
	} else if flag == "sub" {
		return func(x, y int) int {
			return x - y
		}
	} else {
		return nil
	}
}

func main() {
	fmt.Println(getFunc("add")(1, 2))

	fmt.Println(getFunc("sub")(1, 2))

	fmt.Println(getFunc("multiply")(1, 2))
}
