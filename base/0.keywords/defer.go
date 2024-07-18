package main

import (
	"errors"
	"fmt"
)

func f1() {
	var err error
	defer fmt.Println("f1 error :", err) // 作为参数 在defer定义时就会求值 此时即nil
	err = errors.New("defer1 error")
	return
}

func f2() {
	var err error
	defer func() {
		fmt.Println("f2 error :", err) // 闭包
	}()
	err = errors.New("defer2 error")
	return
}

func f3() {
	var err error
	defer func(err error) {
		fmt.Println("f3 error :", err)
	}(err) // 作为参数 在defer定义时就会求值 此时即nil 线上要留心这种情况 避免defer没起作用
	err = errors.New("defer3 error")
	return
}

func main() {
	s := "hello"
	defer fmt.Println(s) // 复制一份`hello`保存起来

	s = "hi"
	fmt.Println(s)

	f1()
	f2()
	f3()
}
