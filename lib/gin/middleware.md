## 前言

如无特殊说明，本文的案例全部来自于本人的一个学习项目。

项目地址：

https://github.com/cauliflower-beep/riceMall.git

## 路由分组

实际开发应用的过程中，我们希望能将各个功能模块的路由进行分组，同一个模块的不同路由带有同样的前缀。这样做的好处是：

1. 路由更加清晰；
2. 针对某一组路由进行中间件权限校验的时候比较方便。

如下所示：

```go
func NewRouter() *gin.Engine {
	r := gin.Default()
	...
	v1 := r.Group("api/v1")
	{
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
        ...
		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)
			authed.POST("avatar", api.UploadAvatar) // 上传头像
			...
		}
	}
	return r
}
```

其中的authed分组使用了自定义的JWT中间件，如果token校验失败，则无法执行对应的handler过程。

## 中间件

### 1.中间件简介

gin允许开发者在处理请求的过程中，加入自己的钩子（Hook）函数。这个钩子函数就叫中间件。

中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。

![中间件](.\img\中间件.png)

如图所示。中间件就是作用于两个模块之间的功能软件，可以作为拦截器、记录日志等。比如在前后端开发中，遵循如下流程：

路由——>中间件(过滤作用)——>控制器

在gin中，中间件的效果可以简单概括为：

1. 设置好中间件之后，后续的路由都会使用这个中间件；
2. 设置在中间件之前的路由则不会生效。

### 2. 中间件原理

上文提到，中间件是为了过滤路由而发明的一种机制，也就是http请求来到时，先经过中间件，再到具体的处理函数。

我们来看一下gin.Context结构体的主要字段:

```go
// gin.Context 结构体
type Context struct {
	...
	handlers HandlersChain 		// 函数指针切片对象
	index 	 int8				// 对应函数指针切片中的索引下标，执行c.Next()时会向后移动index下标位置
	...
}

type HandlerFunc func(*Context)  // 函数指针
type HandlersChain []HandlerFunc // 函数指针切片
```

可以看到，gin框架的中间件函数和处理函数，是以切片形式的调用链条存在的（本质上就是函数指针切片）。

在我们初始化了gin对象之后：r := gin.New()，到中间件生效，往往要经历如下两个过程：

- 注册

  r.Use().也就是不断的在上述的HandlerChain函数指针切片后执行append操作，去依次向调用链条中追加新注册的中间件函数

  ```go
  func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
  	group.Handlers = append(group.Handlers, middleware...)
  	return group.returnObj()
  }
  ```

- 调用

  如下图，gin框架正是通过移动切片下标index的位置，实现中间件的不断向后调用。

  但是这个index要如何去移动呢?下文提到的c.Next()会给我们答案。这个函数在每次调用时，将index向后移动，从而依次调用已注册的中间件。

  ![index](.\img\index.png)

  ```go
  func (c *Context) Next() {
  	c.index++
  	for c.index < int8(len(c.handlers)) {
  		c.handlers[c.index](c)
  		c.index++
  	}
  }
  ```

- 停止调用后续

  c.Abort()方法会阻止调用后续的中间件处理函数，正是因为它使得index移动到切片末尾了，所以后面的`中间件/路由处理函数`都没办法继续执行了。

  ```go
  func (c *Context) Abort() {
  	c.index = abortIndex
  }
  ```

### 3.定义中间件

gin的中间件必须是一个gin.HandlerFunc类型，在自定义中间件函数时，通常采用如下两种写法：

```go
//JWT 中间件-token验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

//JWT 中间件-token验证
func JWT(c *gin.Context) {
    code := 200
    var data interface{}
    token := c.GetHeader("Authorization")
    if token == "" {
        code = 404
    } else {
        claims, err := util.ParseToken(token)
        if err != nil {
            code = e.ErrorAuthCheckTokenFail
        } else if time.Now().Unix() > claims.ExpiresAt {
            code = e.ErrorAuthCheckTokenTimeout
        }
    }
    if code != e.SUCCESS {
        c.JSON(200, gin.H{
            "status": code,
            "msg":    e.GetMsg(code),
            "data":   data,
        })
        c.Abort()
        return
    }
    c.Next()
}

// 路由组添加中间件
authed.Use(middleware.JWT())

```

### 4.注册中间件

gin支持注册全局中间件，也可以给单独路由或者路由组注册中间件。

同时，可以为路由添加任意数量的中间件。

以某个路由为例，当存在多个中间件的时候，处理顺序是参考洋葱模型:

![洋葱模型](.\img\洋葱模型.png)

简而言之，请求是队列处理，响应则是堆栈处理。

#### 1.注册全局中间件

```go
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 增设全局中间件
	r.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	...
	v1 := r.Group("api/v1")
	{
		// 后台接口测试
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		...
	}
	return r
}
```

本地启动项目，在中间件Cors内部打个断点，会发现所有请求打过来，都会进入断点。说明本示例注册了一个全局的中间件。

#### 2.单独注册某个路由中间件

```go
func main() {
	r := gin.Default()
	r.GET("/test", middleW(), func(c *gin.Context) {
		fmt.Println("我在方法内部")
		c.JSON(200, gin.H{
			"msg": "成功了",
		})
	})
	r.Run(":8080")
}

//声明一个中间件
func middleW() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前")
		c.Next()
		fmt.Println("我在方法后")
	}
}
```

这个自行测试。

#### 3.注册路由组中间件

代码示参见第二节，定义中间件部分。他将对一整个路由组生效。

### 4.中间件的嵌套

中间件是可以嵌套的，来认识3个gin中用于中间件嵌套相关的函数：

#### 1. Next()

表示跳过当前中间件剩余内容，执行下一个中间件。当所有操作执行完之后，以出战的执行顺序返回，执行中间件的剩余代码。

```go
//定义中间件1
func middlewOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是1")
		c.Next()
		fmt.Println("我在方法后,我是1")
	}
}

//定义中间件2
func middlewTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是2")
		c.Next()
		fmt.Println("我在方法后,我是2")
	}
}

func main() {
	r := gin.Default()
	//使用多个中间件
	r.GET("/test", middlewOne(), middlewTwo(), func(c *gin.Context) {
		fmt.Println("我在方法内部")
		c.JSON(200, gin.H{
			"msg": "这里是test1",
		})
	})
	r.Run()
}
```

#### 2. return

终止执行当前中间件剩余内容，执行下一个中间件。

当所有的函数执行结束后，以出栈的顺序执行返回，但不执行return后的代码。

```go
//定义中间件1
func middlewOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是1")
		c.Next()
		fmt.Println("我在方法后,我是1")
	}
}

//定义中间件2
func middlewTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是2")
		return
		fmt.Println("我在方法后,我是2")
	}
}

//定义中间件3
func middlewThree() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是3")
		c.Next()
		fmt.Println("我在方法后,我是3")
	}
}
func main() {
	r := gin.Default()
	//使用多个中间件
	r.GET("/test", middlewOne(), middlewTwo(), middlewThree(), func(c *gin.Context) {
		fmt.Println("我在方法内部")
		c.JSON(200, gin.H{
			"msg": "这里是test1",
		})
	})
	r.Run()
}
```

#### 3. Abort()

只执行当前中间件，操作完成后，以出栈的顺序，依次返回上一级中间件。

```go
//定义中间件1
func middlewOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是1")
		c.Next()
		fmt.Println("我在方法后,我是1")
	}
}

//定义中间件2
func middlewTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是2")
		c.Abort()
		fmt.Println("我在方法后,我是2")
	}
}

//定义中间件3
func middlewThree() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是3")
		c.Next()
		fmt.Println("我在方法后,我是3")
	}
}
func main() {
	r := gin.Default()
	//使用多个中间件
	r.GET("/test", middlewOne(), middlewTwo(), middlewThree(), func(c *gin.Context) {
		fmt.Println("我在方法内部")
		c.JSON(200, gin.H{
			"msg": "这里是test1",
		})
	})
	r.Run()
}
```

### 5.注意事项

#### 5.1 gin默认中间件

gin.Default()默认使用了Logger和Recovery中间件，其中：

Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。

Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。

如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

#### 5.2 gin中间件中使用goroutine

当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy())。

#### 5.3 参数传递

借助c.Set()和c.Get()方法，能使我们在不同的中间件函数中传递数据。