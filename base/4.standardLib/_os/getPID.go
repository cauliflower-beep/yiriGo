package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("当前进程ID：", os.Getpid())

	procArr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}

	// 运行计算机上的其他任何程序，包括自身的命令行、java程序等
	process, err := os.StartProcess("./hello", []string{"", "hello,world!"}, &procArr)
	if err != nil {
		fmt.Println("进程启动失败：", err)
		os.Exit(2)
	} else {
		fmt.Println("紫禁城ID：", process.Pid)
	}

	time.Sleep(time.Second)
}
