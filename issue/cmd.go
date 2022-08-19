package main

/*
go1.17 之后，使用go get安装依赖包可能会导致以下告警:
	go get: installing executables with ‘go get‘ in module mode is deprecated.

原因分析：
go1.17 版使用go install安装依赖
	go get用于下载并安装go包、命令等，而go install 在module时代几乎很少使用。
在go path时代，go install 用来编译安装本地项目。
	1.16 时，官方说，不应该使用go get下载安装命令（即可执行程序），不过只是
这样说，依然可以正常使用。
	但自go 1.17 开始，如果继续使用 go get 安装命令会告警。也就是说，go get只用
来下载普通的包，安装可执行程序，应该使用go install。例如：
	go install github.com/github/hub
会将hub命令安装到 $GOBIN 下.此外，go get 有一个参数 -d , 指示 go get 下载对应
的包，但不做编译和安装。未来版本 -d 也会成为默认行为，这样会更快。因为不编译，即使目
标依赖在特定平台编译报错，go get也能正常执行完。

解决方案:
安装远程依赖命令：
	# go 1.16及以前：
	go get -u -v github.com/coocood/freecache

	# go 1.17版本
	go install github.com/coocood/freecache

参考链接：
https://mp.weixin.qq.com/s?__biz=MzAxNzY0NDE3NA==&mid=2247485794&idx=1&sn=092e3d66615680692dd22f5e2f727989&chksm=9be32683ac94af9594b35e9dbf10adc1615110a1be44d39414d29238ed0ca9c1f64a5cb3f814&scene=126&&sessionid=1642213612#rd
*/
