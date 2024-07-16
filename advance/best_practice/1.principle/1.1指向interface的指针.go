package main

import "fmt"

type Foo interface {
	Say(name string)
}

type Dog struct {
	name string
}

func (d Dog) Say(name string) {
	d.name = name
	fmt.Printf("wong!my name is %s\n", d.name)
}

type Cat struct {
	name string
}

func (c *Cat) Say(name string) {
	c.name = name
	fmt.Printf("meow!my name is %s\n", c.name)
}

/*
	f1.Say() 无法修改底层数据
	f2.Say() 可以修改底层数据，给接口变量 f2 赋值时使用的是对象指针
	不需要指向接口类型的指针 直接将接口作为值传递即可 事实上传递的底层数据仍然可以是指针
*/
var f1 Foo = Dog{name: "小白"}
var f2 Foo = &Cat{name: "加菲"}

func main() {
	fmt.Println("f1修改之前：", f1)
	f1.Say("史努比")
	fmt.Println("f1修改之后：", f1)
	fmt.Println(f1.(Dog).name) // 将 f1 接口断言成 Dog{}.即可调用 name 字段
	fmt.Println("______________________")
	fmt.Println("f2修改之前：", f2)
	f2.Say("汤姆")
	fmt.Println("f2修改之后：", f2)
	fmt.Println(f2.(*Cat).name) // 将 f2 接口断言成 *Cat{}.即可调用 name 字段
}
