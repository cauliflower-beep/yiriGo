package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

/*
日常在业务开发中，都希望能有一个trace_id能够串联所有的日志，这就需要我们打印日志时能够获取到这个trace_id.

实现:可以通过 withValue 创建一个携带trace_id的context，然后不断透传下去，从中派生的context都会获取此值，
打印日志时输出即可。

目前一些RPC框架（如rpcx都是支持了Context,trace_id的向下传递就更方便了）

使用withValue时要注意4个事项：
	1.不建议使用context值传递关键参数，关键参数应该显示的声明出来，不应该隐式处理，context中最好是携带签名、trace_id这类值；
	2.因为携带value也是key、value的形式，为了避免因多个包同时使用context而带来冲突（有些包可能获取不到其他包中的自定义类型？），
	  key建议采用内置类型；
	3.本案例中，trace_id是直接从当前的ctx中获取的，实际上也可以获取父context中的value.在获取键值对时，我们先从当前context中
      查找，没找到会从父context中查找该键值对对应的值，直到在某个父context中返回nil或者找到对应的值；
	4.context传递的数据中key、value都是interface类型，这种类型编译期无法确定类型，所以不是很安全，在类型断言时别忘了保证程序的
	  健壮性。
*/

const (
	KEY         = "trace_id"
	DATE_FORMAT = "2006-01-02 15:04:05"
)

func NewRequestID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1) // 用""替换所有的"-"
}

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID()) // 创建一个上下文，将key与uuid绑定
	return ctx
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format(DATE_FORMAT), GetContextValue(ctx, KEY), message)
}

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string) // 取出上下文中key的value
	if ok {
		return v
	} else {
		return ""
	}
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "杀不死我的，只会让我更强大！")
}

func main() {
	ProcessEnter(NewContextWithTraceID())
}
