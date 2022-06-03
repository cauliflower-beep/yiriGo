package main

import (
	"context"
	"fmt"
	"time"
)

/*
通常健壮的程序都是要设置超时时间的，避免因为服务端长时间响应消耗资源，所以一些web框架或者rpc框架都会采用withTimeout或者
withDeadline来做超时控制。当一个请求到达设置的超时时间，就会及时取消，不再往下执行。

withTimeout和withDeadline作用是一样的，只是传递的参数不同。他们都会通过传入的时间自动取消context，这里要注意的是他们
都会返回一个cancelFunc的方法，通过调用这个方法可以提前取消运行。不过在使用的过程还是建议在自动取消后也调用cancelFunc去
停止定制，减少不必要的资源浪费。这两个方法使用哪个都是一样的，看业务场景和个人习惯。本质withTimeout内部也是调用的withDeadline

withTimeout将持续时间作为参数输入。
*/

// NewContextWithTimeout 创建超时控制上下文
func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

/*
到达终止时间之后，自动终止接下来的执行
*/
func dealAutoCancel(ctx context.Context) {
	for i := 1; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done(): // 必须监听取消信号，否则即使传递了ctx,也不会自动终止执行
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("deal time is:", i)
		}
	}
}

/*
未到达终止时间，手动终止接下来的执行
*/
func dealManuCancel(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return // 注意return，否则ctx终止之后会继续执行循环
		default:
			fmt.Println("finish deal")
			cancel()
		}
	}
}
func httphandler() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel()
	/*
		注意：

	*/
	//dealAutoCancel(ctx) // auto cancel
	dealManuCancel(ctx, cancel) //  manual cancel
}

func main() {
	httphandler()
}

/*
https://blog.csdn.net/qq_39397165/article/details/121092507
*/
