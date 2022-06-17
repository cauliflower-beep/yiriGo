package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
)

var (
	addr = flag.String("addr", "localhost:8973", "server address")
)

type Args struct {
	A int
	B int
}
type Rep struct {
	C int
}

func main() {
	flag.Parse() // 解析命令行参数
	// 点对点，客户端直连服务器来获取服务地址
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	// failmode告诉客户端如何处理调用失败，selectmode告诉客户端如何在多台服务器提供了同一服务的情况下选择服务器
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// 定义请求
	args := &Args{
		A: 3,
		B: 6,
	}

	// 定义响应对象，默认0值，rpcx会通过它来知晓返回结果的类型，并把结果反序列化到这个对象
	rep := &Rep{}

	// 1.同步调用远程服务获取结果
	err := xclient.Call(context.Background(), "Mul", args, rep)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, rep.C)

	// 2.异步调用服务,使用 xclient.Go 替换 xclient.Call, 然后把结果返回到一个channel中，可以从channel中监听调用结果
	call, err := xclient.Go(context.Background(), "Mul", args, rep, nil)
	if err != nil {
		log.Fatalf("failed to call： %v", err)
	}
	repCall := <-call.Done
	if repCall.Error != nil {
		log.Fatalf("failed to call： %v", repCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, rep.C)
	}

	// 调用纯函数注册的服务
	// 一个xclient对应一个服务，如果想要调用多个服务，需要为每个服务创建一个xclient
	xclient2 := client.NewXClient("purefunc.service", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient2.Close()
	err = xclient2.Call(context.Background(), "sub", args, rep)
	if err != nil {
		log.Fatalf("failed to call sub: %v", err)
	}
	log.Printf("%d - %d = %d", args.A, args.B, rep.C)
}
