## go protoc-gen-go 安装详解

首先是安装官方的protoc工具，可以从其GitHub官方网站下载

https://github.com/protocolbuffers/protobuf/releases

我下载的win-64版本。

下载完成后解压把bin下面的protoc.exe 拷贝到GOPATH下面。

例如我解压到了：`D:\Go\bin`这个目录。

然后是安装针对Go语言的代码生成插件

```go
go get github.com/golang/protobuf/protoc-gen-go
```

报错如下错误：

```go
go: module github.com/golang/protobuf is deprecated: Use the "google.golang.org/protobuf" module instead.
```

我使用的go版本是1.19。而Go1.17版之后使用go install安装依赖。所以应该按照它下面的格式：

`go install pkg@version`

进行拉取，同时按照提示原来的地址作废，需要替换成google.golang.org/protobuf，最终命令为：

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`