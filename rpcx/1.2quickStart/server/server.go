package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
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

func main() {
	s := server.NewServer() // 创建服务器实例
	// 注册服务
	if err := s.RegisterName("Arith", new(Arith), ""); err != nil {
		fmt.Println("register service failed!")
	}
	// 创建监听
	err := s.Serve("tcp", ":8972")
	if err != nil {
		fmt.Println("server start failed!")
	}

}
