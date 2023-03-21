package main

import (
	"fmt"
	"os"
)

// 判断所给路径是否为文件夹

func isDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func isFile(path string) bool {
	return !isDir(path)
}

func main() {
	dir := isDir("./test")
	if dir {
		fmt.Println("./test 是个文件夹")
	}
	file := isFile("./test")
	if file {
		fmt.Println("./test 是个文件")
	}
}
