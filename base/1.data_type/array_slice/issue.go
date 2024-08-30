package a_s

/*
下面的代码输出什么？
*/

// func main() {
// 	var p [100]int
// 	var m interface{} = [...]int{99: 0}
// 	fmt.Println(p == m)
// 	fmt.Println(p, reflect.TypeOf(p), reflect.ValueOf(p))
// 	fmt.Println(m, reflect.TypeOf(m), reflect.ValueOf(m))
// }

/*
答案：true 二者均为长度100，且元素值全是0的数组

解析：
	1. interface类型变量与非interface类型变量判等时，首先要求非interface类型实现了该接口，否则编译不过。本例中，
	   接口方法集为空，可以认为所有类型都实现了该接口，故可以编译通过；
	2. 满足上一条的前提下，interface类型变量的动态类型、值均与非interface类型变量相同时，两个变量判等结果为true，
	   结合array判等规则，答案为true
*/

/*
通过var s []int 与 s := make([]int,0)定义的切片有什么区别？
*/
