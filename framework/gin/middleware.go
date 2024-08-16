package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middleWOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是1")
		c.Next()
		fmt.Println("我在方法后,我是1")
	}
}

func middleWTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前,我是2")
		c.Next()
		fmt.Println("我在方法后,我是2")
	}
}

func main() {
	r := gin.Default()
	//使用多个中间件
	r.GET("/next", middleWOne(), middleWTwo(), func(c *gin.Context) {
		fmt.Println("我在方法内部")
		c.JSON(200, gin.H{
			"msg": "这里是next",
		})
	})
	_ = r.Run()
}
