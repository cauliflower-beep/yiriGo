## 返回值

golang中的函数可以返回多个返回值，需遵循如下规则：

1. 只要有一个返回值有指定命名，其他的也必须有命名；
2. 如果有多个返回值，必须加括号；
3. 如果只有一个返回值并且有命名，也需要加上括号。

如下代码会编译失败，因为多个返回值有些指定命名有些没有：

```go
func add(x,y int)(sum int,error){
    return x+y,nil
}
...
```

## 递归

使用递归要注意栈溢出的问题。一旦栈溢出，治标的办法是增大栈大小。

治本的办法是修改代码。一般都是代码没写好。

## 函数式编程

golang支持函数式编程。

### 方法做为变量

golang支持定义函数，然后将函数赋值给一个变量，通过变量来调。

为什么不直接调？

主要是支持高度自定义。

### 局部方法

可以再方法内部声明一个局部方法，它的作用域就在本方法内。





