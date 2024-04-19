## 背景

​	在 Windows 环境开发时，有时候需要编译成其他平台（例如Linux）的可执行文件，golang是支持跨平台编译的，只需做相关设置即可：

```go
go env -w CGO_ENABLED = 0
go env -w GOARCH=amd64
// go env -w GOARCH=arm
go env -w GOOS=linux
// go env -w GOOS=windows 
// go env -w GOOS=darwin // mac
go build xxx
```

## goland配置

### 1.输出位置

比如这里设置为：输出到当前目录。

![image-20220817182406617](.\img\image-20220817182406617.png)

### 2.编译配置

![image-20220817183325243](.\img\image-20220817183325243.png)

### 3.编译

![image-20220817184235639](\img\image-20220817184235639.png)

即可在输出目录看到编译生成的可执行文件。

### ps.

linux下生成的可执行文件可能需要先赋予权限再运行：

```shell
chmod 777 crossCompile_linux
```



