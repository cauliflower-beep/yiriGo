package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

/*
	在 net_test.go 这个测试文件中，要写的内容是很多的
	其实 golang 提供了 httptest 这个包来高效测试网络连接

	运行 go test -run TestConn2 查看测试结果
*/

func TestConn2(t *testing.T) {
	// 使用 httptest 模拟请求对象(req)和响应对象(w),达到了相同的目的
	req := httptest.NewRequest("GET", "http://127.0.0.1:9999/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world!" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}
