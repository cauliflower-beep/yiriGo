package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("before all tests")
}

func teardown() {
	fmt.Println("after all tests")
}

/*
	在这个测试文件中，包含有 2 个测试用例 Test1 和 Test2
	如果测试文件中包含函数 TestMain,那么生成的测试将调用TestMain(m),而不是直接运行测试
	调用 m.Run() 触发所有测试用例的执行,并使用 os.Exit() 处理返回的状态码，如果不为0,说明有用例失败
	因此可以在调用 m.Run() 前后做一些额外的准备 (setup) 和回收 (teardown)工作。

	运行	go test -v .\setup_test.go 查看结果
*/
func Test1(t *testing.T) {
	fmt.Println("im test1")
}

func Test2(t *testing.T) {
	fmt.Println("im test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
