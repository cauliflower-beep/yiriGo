## 如何确定是否发生逃逸？

小节代码：[check.go](.\check.go)

### -gcflags

Go 提供了相关的命令，可以查看变量是否发生逃逸。执行如下命令：

```go
go build -gcflags '-m -l' check.go
```

其中 `-gcflags` 参数用于启用编译器支持的额外标志。例如，-m 用于输出编译器的优化细节 （包括使用逃逸分析），相反可以使用 -N 来关闭编译器优化；而 -l 则用于禁用 foo 函数 的内联优化，防止逃逸被编译器通过内联彻底的抹除。

执行后得到如下输出：

```bash
# command-line-arguments
./check.go:6:2: moved to heap: t
./check.go:12:13: ... argument does not escape
./check.go:12:14: *x escapes to heap
```

foo 函数里的变量 t 逃逸了，因为它在函数之外被main函数引用了。

但为什么 main 函数里的 x 也逃逸了？ 这是因为有些函数的参数为 interface 类型，比如 `fmt.Println(a ...interface{})`，编译期间很难确定其参数的具体类型，也会发生逃逸。

### go tool

使用反汇编命令也可以看出变量是否发生逃逸。执行命令：

```bash
go tool compile -S check.go
```

结果中如果包含`runtime.newobject`，说明有变量被放在了堆上，也就是发生了逃逸。因为这个函数用于在堆上分配一块儿内存。