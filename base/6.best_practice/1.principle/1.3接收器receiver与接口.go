package main

import "fmt"

type tiger struct{}

func (t tiger) hunt() {
	fmt.Println("猎个痛快！")
}

func (t *tiger) eat() {
	fmt.Println("一日三餐谁都不能离~")
}

/*
	拓展:
	如果某个结构或类型实现了接口方法，即使方法是值接收器，同样可以用指针接收器来满足接口
	反过来，实体类型以指针接收者实现接口的时候，只有指向这个类型的指针才被认为实现了该接口
	原因:
	接收者是指针类型的方法，很可能在方法中会对接收者的属性进行更改操作，从而影响接收者；
	而对于接收者是值类型的方法，在方法中不会对接收者本身产生影响。
	所以，当实现了一个接收者是值类型的方法，就可以自动生成一个接收者是对应指针类型的方法，因为两者都不会影响接收者。
	但是，当实现了一个接收者是指针类型的方法，如果此时自动生成一个接收者是值类型的方法，原本期望对接收者的改变（通过指针实现），现在无法实现，因为值类型会产生一个拷贝，不会真正影响调用者。
*/
type animal interface {
	run()
}

func (t tiger) run() {
	fmt.Println("如果前面有个悬崖，你早就起飞啦!不是往这儿跑哒小宝贝")
}

type lion struct{}

func (l *lion) run() {
	fmt.Println("逮虾户!")
}
func main() {
	var t1 = tiger{}
	fmt.Println(t1)
	t1.hunt() // 通过值只能调用hunt
	t1.eat()  // 无法编译通过 ? 为啥我这样可以调用成功? todo
	/*
		上述问题似乎找到了答案:某些情况下，值类型也是可以调用指针接收者方法的.
		当类型和方法的接收者类型不同时(如上述`值类型` 调 `指针接收者`),编译器会做一些操作:
		在值类型调用指针接收者方法时, 实际为(&t1).eat()
		在指针类型调用值接收者方法时, 实际为(*t2).hunt()
		这是类型不同可以相互调用的情况

		类型不同不能相互调用的情况,主要有两种:
		1.值类型不能被寻址:
			如果值类型实体不能被寻址，那么它就不能调用指针接收者方法
			type field struct {
				name string
			}
			func (p *field) pointerMethod() {
				fmt.Println(p.name)
			}
			func NewFiled() field {
				return field{name: "right value struct"}
			}
			func main() {
				NewFiled().valueMethod()
				NewFiled().pointerMethod()
			}
			解释:
			看来编译器首先试着给 NewField() 返回的右值调用 pointer method，出错；然后试图给其插入取地址符，未果，就只能报错了。
			至于左值和右值的区别，大家感兴趣可以自行搜索一下。
			大致来说，最重要区别就是是否可以被寻址，可以被寻址的是左值，既可以出现在赋值号左边也可以出现在右边；
			不可以被寻址的即为右值，比如函数返回值、字面值、常量值等等，只能出现在赋值号右边。
		2.用指针接收者实现接口.
		两种类型都是值类型不能调用指针接收者方法.
	*/
	var t2 = new(tiger)
	fmt.Println(t2)
	t2.hunt()
	t2.eat()

	// 拓展部分
	tVal := tiger{}
	tPtr := &tiger{}
	//lVal := lion{}
	lPtr := &lion{}

	var ani animal
	ani = tVal
	ani.run()
	ani = tPtr
	ani.run()
	//ani = lVal // 无法编译通过 因为 lVal是一个值, 而 lion 的 run 方法中没有使用值接收器
	//ani.run()
	ani = lPtr
	ani.run()

}
