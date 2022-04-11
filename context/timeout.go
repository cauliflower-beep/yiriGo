package main

import (
	"context"
	"fmt"
	"time"
)

func NewContextWithTimeout()(context.Context,context.CancelFunc){
	return context.WithTimeout(context.Background(),3 * time.Second)
}

/*
到达终止时间之后，自动终止接下来的执行
 */
func dealAutoCancel(ctx context.Context){
	for i := 1;i <10;i++{
		time.Sleep(1*time.Second)
		select{
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("deal time is:",i)
		}
	}
}

/*
未到达终止时间，手动终止接下来的执行
 */
func dealManuCancel(ctx context.Context,cancel context.CancelFunc){
	for i :=0;i<10;i++{
		time.Sleep(1*time.Second)
		select{
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			return		// 注意return，否则ctx终止之后会继续执行循环
		default:
			fmt.Println("finish deal")
			cancel()
		}
	}
}
func httphandler(){
	ctx , cancel := NewContextWithTimeout()
	defer cancel()
	//dealAutoCancel(ctx)		 // auto cancel
	dealManuCancel(ctx,cancel)	//  manual cancel
}

func main(){
	httphandler()
}


/*
https://blog.csdn.net/qq_39397165/article/details/121092507
 */