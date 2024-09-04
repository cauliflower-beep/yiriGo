package main

import (
	"fmt"
	"runtime"
)

/*
	recover可以让进入宕机流程中的goroutine恢复过来，仅在延迟函数defer中有效
	在正常的执行过程中，调用recover会返回nil并且没有任何其他效果。
	一旦当前goroutine陷入panic,调用recover可以捕获到panic输入值，并且恢复正常的执行

	通常来说，不应该对进入panic宕机的程序做任何处理，但有时，需要我们可以从宕机中恢复，
	或者至少可以在程序崩溃前，做一些操作。
	举个例子，当web服务器遇到不可预料到严重问题时，在崩溃前应该将所有的连接关闭，如果不做任何处理，会使客户端一直处于等待状态
	如果web服务器处于开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试

	在其他语言里，宕机往往以异常的形式存在，底层抛出异常，上层逻辑通过try/cache机制捕获异常，
	没有被捕获的严重异常会导致宕机，捕获的异常可以被忽略，让代码继续运行

	go语言没有异常系统，其使用panic触发宕机，类似于其他语言的抛出异常，recover的宕机恢复机制，对应于其他语言中的try/catch机制
*/

// panicCtx panic时需要传递的上下文信息
type panicCtx struct {
	function string // 所在函数
}

// protectRun 保护方式运行一个函数
func protectRun(entry func()) {
	// 延迟处理函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()
	entry()
}

func main() {
	fmt.Println("运行前...")
	// 运行一段手动触发的错误
	protectRun(func() {
		fmt.Println("手动宕机前...")
		// 使用panic传递上下文
		panic(&panicCtx{
			function: "手动触发panic",
		})
		fmt.Println("手动宕机后...")
	})
	// 故意造成空指针访问错误
	protectRun(func() {
		fmt.Println("赋值宕机前...")
		var a *int
		*a = 1
		fmt.Println(a)
		fmt.Println("赋值宕机后...")
	})
	fmt.Println("运行后...")
}

/*
	panic和recover的组合有如下特性：
	1.有panic没recover，程序宕机；
	2.有panic也有recover，程序不会宕机，执行完对应的defer之后，从`宕机点`退出当前函数后继续执行

	虽然panic/recover能模拟其他语言的异常机制，但并不建议在编写普通函数时也经常性使用这种特性
	在panic触发的defer函数内，可以继续调用panic，进一步将错误外抛，直到程序集体崩溃。
	如果想在捕获错误时设置当前函数的返回值，可以对返回值使用命名返回值方式直接进行设置。
*/
