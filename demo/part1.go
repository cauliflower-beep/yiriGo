package main

import (
	"fmt"
	"net/http"
)

var p = newPerson()

func newPerson() *person {
	return &person{}
}

type person struct{}

func main() {
	rsp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()
	fmt.Println(rsp.Body)
}
