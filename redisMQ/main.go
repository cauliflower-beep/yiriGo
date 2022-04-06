package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/*
思路：利用redis的list做消息队列，左进右出。
 */

const redismq = "redismq"

// Producer
func producer(){
	conn ,err := redis.Dial("tcp","192.168.1.13:6379")
	if err != nil{
		log.Fatal("producer connect failed!",err)
	}
	defer conn.Close()
	rand.Seed(time.Now().Unix())

	i := 1
	for {
		_,err = conn.Do("rpush",redismq,strconv.Itoa(i))
		if err != nil{
			fmt.Println("producer error:",err)
			continue
		}
		fmt.Println("producer element:%d",i)
		time.Sleep(time.Duration(rand.Intn(10))*time.Second)
		i++
	}
}

func consumer(){
	conn ,err := redis.Dial("tcp","192.168.1.13:6379")
	if err != nil{
		log.Fatal("consumer connect failed!",err)
	}
	defer conn.Close()

	for {
		element,err := redis.String(conn.Do("lpop",redismq))
		if err != nil{
			fmt.Println("no msg,sleeping...",err)
			continue
		}else{
			fmt.Println("cosume element:%s",element)
		}
	}
}

func main(){
	//list := os.Args
	//fmt.Println(list)
	//if list[1] == "pro"{
	//	go producer()
	//}else if list[1] == "con"{
	//	go consumer()
	//}
	//go producer()
	go consumer()
	for {
		time.Sleep(time.Duration(10000)* time.Second)
	}
}
