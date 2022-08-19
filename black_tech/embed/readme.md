## 简介

​	特性`//go:embed`,它的作用就是可以在`Go`语言应用程序中包含任何文件、目录的内容。

​	也就是说我们可以把文件以及目录中的内容都打包到生成的`Go`语言应用程序中了，部署的时候，直接扔一个二进制文件就可以了，不用再包含一些静态文件了，因为它们已经被打包到生成的应用程序中了。

## demo

### 目录结构

创建目录：

```shell
mkdir -p embed_example/{static,templates}	#只是示例命令
```

目录结构：

```shell
cd embed_example; tree
.
├── go.mod
├── go.sum
├── hello.txt
├── main.go
├── static
│   └── pics.jpeg
└── templates
    └── index.html

2 directories, 6 files

```

## 代码

见main.go

## 运行

### 编译运行

执行`go build`生成的可执行文件:

```go
go build -o server main.go
```

### 访问

浏览器访问 `http://127.0.0.1:8080` ，请求注册的接口：

- `http://127.0.0.1:8080/static/pics.jpeg` : 图片`pics.jpeg`
- `http://127.0.0.1:8080/` : templates/index.html

工程中的静态文件已经被打包到编译后的可执行文件中了。