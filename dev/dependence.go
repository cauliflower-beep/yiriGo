package practice

import (
	"fmt"
	"os"
)

/*
	在 Go 语言中，有一句谚语指出了 ”复制“ 的有益之处，叫做："A little copying is better than a little dependency"
	即复制一点总比依赖一点好。
*/

/*
	复制，只要核心
	如果可以自己写一些短小精悍的代码，那就没有必要直接导入一个库去做（可以只复制核心算法）。
	例如 UUID 的案例：
*/

func main() {
	f, _ := os.Open("/dev/urandom")
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uuid)
}

/*
	虽然有很多 UUID 的第三方库，但普遍会有许多功能堆积在一个库中，这样会引入许多不必要的新依赖。
	如果只是要一点新功能，可以自己简单实现，封装为公司内部方法导入。
	可以有效减少依赖管理的负担，缩小二进制文件大小，带来更大的稳定性、安全、测试第三方库这方面大多都是不清楚的。

	 Go 的这句谚语 "A little copying is better than a little dependency"，更多的是一种软件工程里的指导思想。
	当你只是涉及到一个很简单的功能，那完全可以自行实现或复制核心代码。没必要直接导入一个大的第三方库.
	它有可能带来许多奇奇怪怪的依赖，使得你的编译构建变得缓慢，依赖管理也复杂了起来。
	这是需要我们都好好思考的。
*/
