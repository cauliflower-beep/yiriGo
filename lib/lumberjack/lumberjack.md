[源码地址](github.com/natefinch/lumberjack)

## 简介

lumberjack提供了一个方便的方式来处理[日志文件](https://so.csdn.net/so/search?q=日志文件&spm=1001.2101.3001.7020)的轮换，以防止日志文件无限增长。

## 特性

它的主要特性有：

- 日志轮换：
  1. 基于日志文件大小。当日志文件大小达到指定的大小限制时，它会自动进行日志轮换，将日志写入一个新的文件中。
  2. 基于日志文件年龄：可以设置日志文件的最大年龄。当日志文件的年龄超过指定的天数时，它也会进行轮换。
- 备份：该库支持保留一定数量的备份日志文件。备份通常以递增的编号命名，例如 yourlog.log、yourlog.log.1、yourlog.log.2 等等。
- 高性能：lumberjack 专为高性能日志记录而设计。它以异步方式写入日志条目，允许应用程序在无需等待日志写入完成的情况下继续运行，从而减少性能影响。

## 使用

1. **导入包**：

   ```go
   import "github.com/natefinch/lumberjack"
   ```

2. **创建 Lumberjack 日志记录器**：

   创建 `lumberjack.Logger` 结构的新实例，指定日志文件的名称、最大大小、最大备份数和最大保存天数。例如：

   ```go
   logger := &lumberjack.Logger{
       Filename: "myapp.log",
   	MaxSize: 100, // 兆字节
   	MaxBackups: 3,
   	MaxAge: 28, // 天数
   }
   ```

   这个实例将负责处理日志文件的轮换和管理。

3. 设置 Go 日志记录器的输出：

   如果使用 Go 的标准 log 包进行日志记录，可以将 lumberjack.Logger 设置为日志记录器的输出。这可以通过以下方式完成：

   ```go
   log.SetOutput(logger)
   ```

   这样，通过 log.Print()、log.Println() 或 log.Printf() 创建的任何日志条目都将写入由 lumberjack 管理的日志文件。

4. **编写日志条目**：

   使用 Go 的标准日志记录函数来编写日志条目。例如：

   ```go
   log.Println("这将被写入由 lumberjack 管理的日志文件。")
   ```

5. **关闭日志记录器**：

   在应用程序退出时，或在适当的时机，请确保关闭 `lumberjack.Logger` 以确保刷新任何剩余的日志条目并正确关闭日志文件。这可以通过以下方式完成：

   ```go
   logger.Close()
   ```

## 示例

如下是一个完整的使用示例：

```go
package main

import (
    "log"
    "github.com/natefinch/lumberjack"
)
func main() {
    logger := &lumberjack.Logger{
        Filename: "myapp.log",
        Max
        Size: 100, // 兆字节
        MaxBackups: 3,
        MaxAge: 28, // 天数
	}
    defer logger.Close()
    log.SetOutput(logger)
    log.Println("这将被写入由 lumberjack 管理的日志文件。")
}
```

## 参考

[Golang 语言三方库 lumberjack 日志切割组件怎么使用？-腾讯云开发者社区-腾讯云](https://cloud.tencent.com/developer/article/1786358)



