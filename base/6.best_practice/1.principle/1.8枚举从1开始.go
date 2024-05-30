package main

import "fmt"

type Operation int

// Bad
//const (
//	Add      Operation = iota // 0
//	Subtract                  // 1
//	Multiply                  // 2
//)

// Good
const (
	Add Operation = iota + 1
	Subtract
	Multiply
)

// 当零值是默认的理想行为时，使用零值是有意义的
type LogOutput int

const (
	LogToStdout LogOutput = iota
	LogToFile
	LogToRemote
)

func main() {
	fmt.Println(Add, Subtract, Multiply)
	fmt.Println(LogToStdout, LogToFile, LogToRemote)
}
