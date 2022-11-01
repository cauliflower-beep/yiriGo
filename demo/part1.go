package main

import (
	"fmt"
	"strings"
)

var p = newPerson()

func newPerson() *person {
	return &person{}
}

type person struct{}

func main() {
	//rsp, err := http.Get("https://www.baidu.com")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer rsp.Body.Close()
	//fmt.Println(rsp.Body)

	canonicalUri := "/riskwarn/stocks/composite-risks/1/10"
	contextPath := strings.Split(canonicalUri, "/")[1]
	fmt.Println(contextPath)
}
