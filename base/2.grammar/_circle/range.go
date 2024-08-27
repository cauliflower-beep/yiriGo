package main

import "fmt"

/*
Go 不提供类似 C 支持的 while、do...while 等循环控制语法，而仅保留了一种语句，即 for 循环.

但是，经典的三段式循环语句，需要获取迭代对象的长度 n.
鉴于此,为了更方便 Go 开发者对复合数据类型(array_slice/slice/channel/map)进行迭代,Go 提供了 for 循环的变体: for range 循环.
*/

// sliceIteration
//
//	@Description: 切片/数组迭代：
//	@Description: 参与迭代的只是对象的副本，且该对象只会运算一次（如果在循环内部修改切片的长度，不会改变本次循环的次数）
//	@Description: 迭代过程中，下标和值被赋值给变量i和s，第二个参数s是可选的
//	@Description: 针对nil切片，迭代次数为0
func sliceIteration() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a) // [1 2 3 4 5]

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	/*
		r为什么不是 [1 12 13 4 5]？
		--	参与 for range 循环是 range 表达式的副本。也就是说，实际上参与循环的是 a 的副本（go 临时分配的连续字节序列，
		  	与 a 不是同一块内存空间），而不是真正的 a。
	*/
	fmt.Println("after for range loop, r =", r) // [1 2 3 4 5]
	fmt.Println("after for range loop, a =", a) //  [1 12 13 4 5]
}

// mapIteration
//
//	@Description: map迭代，与 array_slice/slice 不同：
//	@Description: 迭代过程中，删除还未迭代到的键值对，则该键值对不会被迭代
//	@Description: 迭代过程中，如果创建新的键值对，那么新增的键值对，可能被迭代，也可能不被迭代
//	@Description: 针对 nil map，迭代次数为0
func mapIteration() {
	roles := map[string]string{
		"火影":  "鸣人",
		"名侦探": "柯南",
		"海贼王": "路飞",
		"龙珠":  "卡卡罗特",
	}

	for k, v := range roles {
		delete(roles, "名侦探")
		roles["一拳超人"] = "埼玉"
		fmt.Printf("%v:%v\n", k, v)
	}
}

// chanIteration
//
//	@Description: 信道channel迭代
//	@Description: 发送给channel的值可以使用for迭代，直到信道被关闭
//	@Description: 如果是nil信道，循环将永远阻塞
func chanIteration() {
	ch := make(chan string) // 无缓冲
	go func() {
		ch <- "for"
		ch <- "range"
		ch <- "channel"
		ch <- "迭代规律"
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
func main() {
	sliceIteration()
	mapIteration()
	chanIteration()
}
