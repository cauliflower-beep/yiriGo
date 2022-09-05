package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:9981", "server2 address")
)

type Args struct {
	A int
	B int
}
type Rep struct {
	C int
}

func main() {
	flag.Parse()

	d, _ := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})

	// failmode告诉客户端如何处理调用失败，selectmode告诉客户端如何在多台服务器提供了同一服务的情况下选择服务器
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	// 定义请求
	args := &Args{
		A: 3,
		B: 6,
	}

	for {
		rep := &Rep{}

		err := xclient.Call(context.Background(), "Mul", args, rep)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		log.Printf("%d * %d = %d", args.A, args.B, rep.C)
		time.Sleep(2 * time.Second)
	}

}
