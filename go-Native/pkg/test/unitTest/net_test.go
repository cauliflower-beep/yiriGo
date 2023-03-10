package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	// 监听一个未被占用的端口，并返回 Listener
	ln, err := net.Listen("tcp", "127.0.0.1:9999")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil) // 启动 http 服务

	// 发起一个 Get 请求 并检查返回值是否正确
	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world!!" {
		t.Fatal("expected hello world,but got", string(body))
	}
}
