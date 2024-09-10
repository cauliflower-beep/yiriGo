package _log

import (
	"log"
	"os"
)

/*
log包定义了Logger类型，提供了一些格式化输出的方法
也提供了一个标准(默认)的`logger`实例，可以调用如下系列函数：
1.Print：Print|Printf|Println
2.Fatal：Fatal|Fatalf|Fatalln
3.Panic：Panic|Panicf|Panicln
*/

func logDefault() {
	// 默认情况下，日志将被输出到标准错误输出(os.Stderr)
	log.Println("info msg")
	log.Fatalln("fatal msg")
	log.Panicln("panic msg")
}

/*原生log提供了一些全局变量和函数用于配置输出行为，但总体来说可配置性较低*/

/*
默认情况下logger只会提供日志时间
如果希望记录更多信息，比如文件名及行号，需要通过setFlags函数定制
*/
func logWithFlags() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	log.Println("info msg")
}

/*
SetPrefix()可以配置日志前缀
*/
func logWithPrefix() {
	log.SetPrefix("[yiri-Go]") // 设置前缀
	log.Println("info msg")
}

/*
原生log无法高度自定义，不能支持格式化输出，也不区分日志级别。
这里再演示一下如何使用原生log将日志内容输出到文件中
*/
func log2File() {
	// 设置日志输出的位置
	file, err := os.OpenFile("log1.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer file.Close()

	log.SetOutput(file)

	// 写入日志消息
	log.Println("info msg") // 从日志文件中可以看到，log包默认使用时间戳和日志内容进行日志记录
}

/*以上演示的是默认Logger，log包也支持定义Logger实例*/

// https://segmentfault.com/a/1190000040977469
