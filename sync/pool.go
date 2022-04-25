package main

import (
	"fmt"
	"time"
)

/*
sync.Pool用很少的代码实现了很巧的功能.
见字如面,Pool一眼就容易联想到池子,而元素池化是常用的性能优化手段(性能优化的几把斧头：并发、预处理、缓存).
比如，创建一个100个元素的池，然后就可以在池子里直接获取到元素，免去了申请和初始化的流程，大大提高了性能.
释放元素也是直接丢回池子,而免去了真正释放元素带来的开销.
sync.Pool除了最常见的池化提升性能的思路，最重要的时减少GC.
它常用于一些对象实例创建昂贵的场景.
注意！Pool是Goroutine并发安全的
 */

func GetNowTimeStr() string {
	nowTime := time.Now()
	// return fmt.Sprintf("%02d:%02d %02d-%02d-%04d", nowTime.Minute(), nowTime.Hour(), nowTime.Day(), nowTime.Month(), nowTime.Year())
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d ", nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute())
}

var uid = 123564898
func main(){
	fmt.Println(GetNowTimeStr())
	//fmt.Printf("%t",string(uid))
	fmt.Println(18.5*0.2)
}
