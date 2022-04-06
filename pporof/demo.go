package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main(){
	ip:= "172.30.42.178:6060"
	fmt.Println("pprof watching...")
	if err := http.ListenAndServe(ip,nil);err != nil{
		fmt.Printf("start pprof failed on %s \n",ip)
	}
}
