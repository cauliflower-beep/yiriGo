package _log

import "log"

/*
log包定义了Logger类型，提供了一些格式化输出的方法
也提供了一个标准的`logger`实例，可以调用如下系列函数：
1.Print：Print|Printf|Println
2.Fatal：Fatal|Fatalf|Fatalln
3.Panic：Panic|Panicf|Panicln
*/

func logDefault() {
	log.Println("info msg")
	log.Fatalln("fatal msg")
	log.Panicln("panic msg")
}

/*
默认情况下logger指挥提供日志时间
如果希望记录更多信息，比如文件名及行号，需要通过setFlags函数定制
*/
func logWithFlags() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	log.Println("info msg")
}

// https://segmentfault.com/a/1190000040977469
