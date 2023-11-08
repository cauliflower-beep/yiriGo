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

Go1.17版之后使用go install安装依赖。所以应该按照它下面的格式：

`go install pkg@version`

进行拉取，同时按照提示原来的地址作废，需要替换成google.golang.org/protobuf，最终命令为：

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

如果要生成grpc代码，需要安装以下插件：
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

进入proto文件所在目录，执行如下生成命令：
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello.proto`
即可生成 .pb.go 与 _grpc.pb.go 文件。

任何更新内容可以参考官网文档：
`https://grpc.io/docs/languages/go/quickstart/`
