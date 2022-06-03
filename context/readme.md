## 作用

context是用来做并发控制的。作用就是在不同的goroutine之间同步请求特定的数据、取消信号以及处理请求的截止日期。

目前我们常用的一些库都是支持context的，例如 gin、database/sql 等库都是支持context的，这样更方便做并发控制，只需要在服务器入口创建一个context上下文，不断透传下去即可。