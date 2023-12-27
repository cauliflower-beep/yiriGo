package main

import (
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = client.Call(
		"KVStoreService.Set", [2]string{"蜡笔小新主角", "风间"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
