package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

/*
	cobra既是一个用来创建现代CLI(Command Line Interface)命令行工具的golang库，也是一个生成程序应用和命令行文件的程序。
	使用 Linux 系统的大部分用户更习惯采用 CLI 进行人机交互。在Win系统诞生以前，它的前身DOS系统也是字符界面的 CLI 人机交互系统;
	而目前Win系统的程序大部分是图形界面GUI(Graphical User Interface)程序，但由于CLI的一些特点和独特的优势，它还不能被完全抛弃。
*/
func main() {
	var rootCmd = &cobra.Command{ // 定义CLI工具的根命令
		// 根命令配置
		Use:   "demo",                                    // 短的单词命令，用于调用CLI工具
		Short: "A brief description of your application", // --help中显示的简要描述
		Long: `A longer description that spans multiple lines 
				and likely contains examples and usage of using your application.`, // --help中显示的详细描述
		Run: func(cmd *cobra.Command, args []string) { // 调用命令时执行的函数
			fmt.Println("Hello, World!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
