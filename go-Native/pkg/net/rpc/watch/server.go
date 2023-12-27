package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type KVStoreService struct {
	// m 用于存储 KV 数据
	m map[string]string
	// filter 对应每个Watch调用时定义的过滤器函数列表
	filter map[string]func(key string)
	// 多个Goroutine并发访问m时的互斥锁
	mu sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered

	p.mu.Lock()
	// 过滤器函数非自定义 当有key变动时就把key写入ch
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}

func main() {
	_ = rpc.RegisterName("KVStoreService", NewKVStoreService())

	// 建立 TCP 连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	var connNum int
	for {
		// 开始监听 如果有请求过来，则返回建立的TCP连接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		connNum++
		fmt.Printf("tcp链接已建立，即将开启rpc服务|%d\n", connNum)

		// 通过 rpc.ServeConn 函数在 TCP 连接上为对方提供 RPC 服务。
		go rpc.ServeConn(conn)
	}
}
