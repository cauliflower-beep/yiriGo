

## 可视化

go语言提供了两种方式可以查看一个程序的GMP数据。

### go tool trace

trace记录了运⾏时的信息，并提供可视化的Web⻚⾯。

小节代码：[trace](.\trace.go).

运行之后会在当前路径下得到一个trace.out文件，可以用工具go tool打开，结果如下：
```go
PS xx\GMP> go tool trace .\trace.out
2024/07/16 18:41:59 Parsing trace...
2024/07/16 18:41:59 Splitting trace...
2024/07/16 18:41:59 Opening browser. Trace viewer is listening on http://127.0.0.1:8912
```

浏览器打开http://127.0.0.1:8912，单击[view trace](http://127.0.0.1:8912/trace)能够看见调度流程：

![](.\trace.png)

信息字段说明：

- Goroutines：G协程信息；
- Heap：堆栈信息；
- Threads：M线程信息；
- ProcX：P调度器信息；

## reference

《深入理解Go语言》- 刘丹冰

