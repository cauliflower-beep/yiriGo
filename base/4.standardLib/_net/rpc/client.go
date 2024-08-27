package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 通过 rpc.Dial 拨号 RPC 服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	// 通过 client.Call 调用具体的 RPC 方法
	err = client.Call("HelloService.Hello", "小新", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
