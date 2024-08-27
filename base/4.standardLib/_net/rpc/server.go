package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	// 将对象类型中所有满足 RPC 规则的对象方法注册为 RPC 函数
	// 注册的方法放在 “HelloService” 服务空间之下
	rpc.RegisterName("HelloService", new(HelloService))

	// 建立 TCP 连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 通过 rpc.ServeConn 函数在 TCP 连接上为对方提供 RPC 服务。
		rpc.ServeConn(conn)
	}
}
