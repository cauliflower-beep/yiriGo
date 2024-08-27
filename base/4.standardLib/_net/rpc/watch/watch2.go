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

	var keyChanged string
	err = client.Call("KVStoreService.Watch", 30, &keyChanged)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("watch:", keyChanged)
}
