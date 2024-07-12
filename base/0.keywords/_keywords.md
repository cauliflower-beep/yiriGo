## select

### 监控case

select 语句用于监控并选择一组case语句，并执行相应的代码。

它看起来类似于 switch 语句，但是所有 case 中的表达式都必须是 channel 的发送或接收操作：

```go
...
select {
	case <-ch1:
		fmt.Println("ch1 ...")
		break
	case <-ch2:
		fmt.Println("ch2 ...")
		break
	default:
		fmt.Println("default process...")
	}
...
```

### deadlock

先看一段代码：

```go
```

