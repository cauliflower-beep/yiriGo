package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:8973", "server2 address")
)

type Args struct {
	A int
	B int
}
type Rep struct {
	C int
}

// 定义一个服务
type Arith struct {
}

// 为服务实现一个方法
func (arth *Arith) Mul(ctx context.Context, args *Args, rep *Rep) error {
	rep.C = args.A * args.B
	return nil
}

func creatServer(addr string) {
	s := server.NewServer() // 创建服务器实例
	// 注册服务
	if err := s.RegisterName("Arith", new(Arith), ""); err != nil {
		fmt.Println("register service failed!")
	}

	// 创建监听
	err := s.Serve("tcp", addr)
	if err != nil {
		fmt.Println("server start failed!")
	}
}
func main() {
	flag.Parse()

	// 启动两个相同的服务，在不同端口监听
	go creatServer(*addr1)
	go creatServer(*addr2)

	select {}

}
