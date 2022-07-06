package main

import (
	"encoding/json"
	"sync"
	"testing"
)

// 运行 go test .\sync.Pool_test.go -bench . -benchmem 测试
type Student struct {
	Name string
	Age  int32
	//Remark [1024]byte
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}
var buf, _ = json.Marshal(Student{Name: "luffy", Age: 21})

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		/*
			json反序列化在文本解析和网络通信过程中十分常见，当程序并发很高时，短时间内需要创建大量的临时变量。
			这些对象分配在堆上，会给GC造成很大压力，严重影响程序性能。
		*/
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		//defer studentPool.Put(stu)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

/*
结果说明：
结果项					含义
BenchmarkUnmarshal-12	BenchmarkUnmarshal 是测试的函数名 -12 表示GOMAXPROCS（线程数）的值为12
1666570					表示一共执行了1666570次，即b.N的值
734.2 ns/op				表示平均每次操作花费了734.2纳秒
248 B/op				表示每次操作申请了248 Byte的内存申请
7 allocs/op				表示每次操作申请了7次内存

可以看到，使用sync.Pool 之后，内存占用是比未使用时要小的
*/
