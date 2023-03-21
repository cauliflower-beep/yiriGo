package main

import (
	"fmt"
	"os"
)

// 判断所给路径文件/文件夹是否存在

func existOrNot(path string) bool {
	_, err := os.Stat(path) // 获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func main() {
	isExist := existOrNot("./春晓.txt")
	if isExist {
		fmt.Println("文件存在")
		os.Exit(0)
	}
	fmt.Println("文件不存在")
}
