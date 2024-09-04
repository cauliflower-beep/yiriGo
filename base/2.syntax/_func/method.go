package _func

import "fmt"

/*
方法是具有接收者的函数
*/
type film struct {
	cast  []string // 卡司
	title string   // 电影名
	date  string   // 公映时间
}

func (fm *film) release() {
	fmt.Println("<", fm.title, ">", "将在", fm.date, "全国公映!")
}

/*
方法的接收者是值类型，调用者既可以是对象，也可以是对象指针
但无论是对象还是对象指针，修改的都是对象的副本，不影响调用者
ps. 可以在一个*T类型的变量P上面调用这个method，而不需要 *P(取指针变量P的值)去调用这个method
*/
type bird struct {
	color string
}

func (b bird) changeColor(color string) {
	b.color = color
}

/*
方法的接收者是指针类型，调用者既可以是对象，也可以是对象指针
但无论是对象还是对象指针，修改的都是调用者本身
ps. 不必 &V 去调用(当然这样也可以),直接用 V.method,Go会识别receiver是指针，自动帮转。
*/
func (fm *film) changeTitle(title string) {
	fm.title = title
}

/*
总结:你不用担心你是调用的指针的method还是不是指针的method，Go知道你要做的一切，这对于有多年C/C++编程经验的同学来说，真是解决了一个很大的痛苦。

通常我们使用指针类型作为方法接收者的理由：
使用指针类型能够修改调用者的值；
使用指针类型可以避免在每次调用方法时复制该值，在值的类型为大型结构体时，这样做会更加高效。
*/

/*
2. method 继承
method是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method
*/

/*
3. method重写
如果不想使用匿名字段的method，结构体可以直接实现自己的 method.
*/
