## grpc

## 1.文档

官方文档：https://grpc.io/docs/

中文文档（开源中国）：http://doc.oschina.net/grpc

笔记资料：https://skyao.gitbooks.io/learning-grpc/content/introduction/information.html

## 2.依赖插件

### 2.1protoc

首先是安装官方的`protoc`工具：

https://github.com/protocolbuffers/protobuf/releases

选择自己操作系统对应的版本，例如win-64。解压，并把bin下面的protoc.exe 拷贝到**GOPATH\bin**中，目的是可以全局使用这个工具。打开终端，键入如下命令查看protoc版本：
```bash
protoc --version
```

正常输出极为安装成功。

### 2.2protoc-gen-go

接着安装针对Go语言的代码生成插件：

```go
go get github.com/golang/protobuf/protoc-gen-go
```

报错如下错误：

```go
go: module github.com/golang/protobuf is deprecated: Use the "google.golang.org/protobuf" module instead.
```

Go1.17版之后使用go install安装依赖。所以应该按照`go install pkg@version`格式进行拉取，同时依据提示替换包地址，最终命令为：

```go
// 安装最新版本的插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

同样查看该插件版本：

```bash
protoc-gen-go --version
```

正常输出版本信息即为安装成功。

### 2.3protoc-gen-go-grpc

如果要生成grpc代码，需要安装如下插件：

```go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

查看插件版本：

```bash
protoc-gen-go-grpc --version
```

正常输出版本信息即为安装成功。

### 注意

protoc-gen-go 和 protoc-gen-go-grpc 都是用于生成 Go 代码的 Protocol Buffers 插件，但它们的功能有所不同。

protoc-gen-go生成用于序列化和反序列化 Protobuf 消息的 Go 代码，这些代码可用于任何使用 Protobuf 进行通信的应用程序，包括但不限于 gRPC 应用程序；

而 protoc-gen-go-grpc 生成用于 gRPC 服务器和客户端的 Go 代码，这些代码基于 protoc-gen-go 生成的代码，并添加了用于实现 gRPC 服务和客户端所需的功能。

简而言之，前者 是一个通用的 Protobuf 插件，而 protoc-gen-go-grpc 是一个专门用于 gRPC 的 Protobuf 插件。

以下是一些有关何时使用每个插件的更具体指南：

- 如果需要编写不使用 gRPC 的 Protobuf 应用程序，则只需使用 protoc-gen-go。
- 如果需要编写 gRPC 服务器，则需要使用 protoc-gen-go-grpc。
- 如果需要编写 gRPC 客户端，则可以选择使用 protoc-gen-go-grpc 或 protoc-gen-go。 使用 protoc-gen-go-grpc 将为我们提供更易于使用的代码，但使用 protoc-gen-go 会给您更多的控制权。

以下是如何使用每个插件的示例命令：

**使用 protoc-gen-go：**

```bash
protoc --go_out=plugins=grpc:path/to/output_directory service.proto
```

**使用 protoc-gen-go-grpc：**

```bash
protoc --go_out=plugins=grpc:path/to/output_directory service.proto
```

## 3.一个栗子

进入proto文件所在目录，执行如下生成命令：
```go
// 生成 hello.pb.go 文件
// --go_out已经指定了使用 protoc-gen-go 插件来生成go代码
// 同理，如果要生成python代码，可以指明参数：--python_out，前提要安装python生成插件
protoc --go_out=. --go_opt=paths=source_relative ./hello.proto

// 生成 hello_grpc.pb.go 文件，包含客户端和服务端的一些代码
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello.proto

// 或 同时生成 hello.pb.go 及 hello_grpc.pb.go 文件
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello.proto
```

命令执行之后，即可在当前目录生成 .pb.go 与 _grpc.pb.go 文件。

说明下上述命令中的`--go_opt=paths=source_relative`参数，它表示 `protoc` 使用与输入 Protobuf 文件相同的路径名来命名生成的 Go 文件。 例如，如果输入文件名为 `hello.proto`，则生成的 Go 文件将名为 `hello.pb.go`。当然也可以自定义生成的名称，具体方法自行搜索吧，目前我接触到的项目，都是默认根据proto的文件名生成的。

任何更新内容可以参考[官网文档](https://grpc.io/docs/languages/go/quickstart/)。
