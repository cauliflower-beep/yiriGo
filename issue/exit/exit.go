package main

import (
	"fmt"
	"os"
)

/*
os.Exit()函数用来终止程序;

os.Exit 函数可以让当前程序以给出的状态码 code 退出，一般来说，状态码 0 表示成功，非 0 表示出错。程序会立即终止，并且defer的函数不会执行。

注意与panic的区别，panic可以触发defer延迟语句，panic还可以被recover捕获处理。

与 return 的区别：
1.return 是返回函数值，是关键字，exit是一个函数；
2.return 是语言级别的，他表示了调用堆栈的返回;而exit是系统调用级别的，他表示了一个进程的结束。
3.return 是函数的退出（返回）;exit是进程的退出。
4.return 是c语言提供的，exit是操作系统提供的(或者函数库中给出的)。
exit是一个库函数，exit(1）表示发生错误后退出程序，exit(0)表示正常退出。
在stdlib.h中exit函数是这样子定义的：void exit(int status)。这个系统调用是用来终止一个进程的，无论在程序中的什么位置，只要执行exit，进程就会从终止进程的运行。
讲到exit这个系统调用，就要提及另外一个系统调用，_exit,_exit()函数位于unistd.h中，相比于exit()，_exit()函数的功能最为简单，直接终止进程的运行，释放其所使用的内存空间，并销毁在内存中的数据结构，
而exit()在于在进程退出之前要检查文件的状态，将文件缓冲区中的内容写回文件。
5.return 用于结束一个函数的执行，将函数的执行信息传出给其他调用函数使用;exit函数是退出应用程序，删除进程使用的内存空间；
并将应用程序的一个状态返回给os或其父进程，这个状态标识了应用程序的一些运行信息，这个信息和机器和操作系统有关，一般是0为正常退出，非0为非正常退出。
6.非主函数中调用return和exit效果很明显，但是main函数中调用return和exit的现象就很模糊，多数情况下现象都是一致的。
*/

func main() {
	var num = 0

	fmt.Printf("enter number:")
	_, _ = fmt.Scanf("%d", &num)

	if num > 0 {
		fmt.Println("Program terminated...")
		os.Exit(-1)
	}
	fmt.Println("Program finished normally")
}
