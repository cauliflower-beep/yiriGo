## 编译和运行程序

### go build

编译Go源代码并生成可执行文件。例如，如果Go源代码文件名为 `main.go`，则可以使用以下命令进行编译：

```go
go build main.go
```

命令将在当前目录下生成一个名为 `main` 的可执行文件。

### go run

编译并运行Go源代码文件，例如：

```go
go run main.go
```

命令将生成一个临时可执行文件，并在运行后将其删除，不会再当前目录下留下可执行文件。

参考：https://blog.csdn.net/weixin_52690231/article/details/124015150#4go_run_75

## 管理依赖项

### go get

下载和安装Go依赖项。例如，可以使用以下命令安装 `github.com/gorilla/mux` 包：

```go
go get github.com/gorilla/mux
```

这将下载并安装该包及其依赖项到`gopath/pkg`目录中。

### go install



### go mod

管理Go模块。Go modules 是Go 1.11中引入的依赖项管理系统。可以使用 `go mod` 命令来查看、添加、编辑和删除模块。

注意，在`gopath`外的目录执行`go mod init`必须指定模块名，不然会报错：

```go
go: cannot determine module path for source directory xxx (outside GOPATH, module path must be specified)

Example usage:
        'go mod init example.com/m' to initialize a v0 or v1 module
        'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.
```

在`gopath/src`下的项目根目录中执行这个命令不会有问题。

## 测试程序

### go test

运行Go程序的测试。例如，如果Go程序包含名为 `TestHello` 的测试函数，可以使用以下命令运行测试：

```go
go test
```

这将运行所有以 `Test` 开头的测试函数。

### go vet

检查Go源代码是否存在潜在的错误和问题。例如，可以使用以下命令检查 `main.go` 文件：

```go
go vet main.go
```

这将检查代码是否存在可能导致问题的潜在问题，例如未使用或未定义的变量。

## 其他

### go fmt

格式化Go源代码。例如，可以使用以下命令格式化 `main.go` 文件：

```go
go fmt main.go
```

这将根据Go代码格式化规范格式化代码。

### go doc

查看Go包或类型的文档。例如，可以使用以下命令查看 `fmt` 包的文档：

```go
go doc fmt
```

这将显示有关 `fmt` 包及其函数和类型的文档。

### go env

打印Go环境信息。例如，可以使用以下命令查看Go安装路径：

```go
go env GO安裝路径
```

这将显示Go安装路径。

### go generate

`go generate` 命令是一个用于自动化生成Go代码的工具。

在Go源文件中的**特殊注释中指定命令**，然后在运行 `go generate` 命令时**自动执行这些命令**。这些命令可以用来**生成代码、格式化代码、运行测试**等等。

命令语法：

```go
go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]
```

`go generate`必须**显示执行**，也就是说在go build、go test之前运行该命令，先创建或更新 GO 源文件，再进行后续的运行测试。

注释格式：

```go
//go:generate command argument...
```

在Go源文件中使用 `go generate` 命令需要在源代码中按如上格式添加特殊注释，其中`command` 是要执行的命令， `arguments` 是传递给该命令的参数。

[示例代码](./generate.go)。

## 更多

有关go命令的更多信息，可以参阅[Go官方文档](https://go.dev/doc/)。

