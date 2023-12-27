package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var v string
	err = client.Call("KVStoreService.Get", "蜡笔小新主角", &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key's v|", v)
}
