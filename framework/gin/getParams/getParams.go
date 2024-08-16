package getParams

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取url路径中的参数 e.g path/book?author=毛姆&price=50
func getUrlParams() {
	router := gin.Default()

	router.GET("/book", func(c *gin.Context) {
		author := c.Query("author")
		price := c.DefaultQuery("price", "50")

		c.String(http.StatusOK, "author: %s, price: %s", author, price)
	})
	_ = router.Run(":8080")
}

// 获取form表单中的数据
func getFormData() {
	router := gin.Default()

	router.POST("/book", func(c *gin.Context) {
		ver := c.PostForm("version")
		price := c.DefaultPostForm("price", "50")

		// fmt.Printf("ver: %s; price: %s; author: %s; date: %s", ver, price, author, date)
		c.String(http.StatusOK, "version: %s, price: %s", ver, price)
	})

	_ = router.Run(":8080")
}

// 获取动态路由path中的参数 e.g /book/月亮与六便士
func getDynamicParams() {
	router := gin.Default()

	// 可以匹配诸如 /book/龙珠 但不会匹配 /book 或者 /book/
	router.GET("/book/:title", func(c *gin.Context) {
		title := c.Param("title")
		c.String(http.StatusOK, "title: %s", title)
	})

	// 可以匹配诸如 /book/龙珠/鸟山明 或者 /book/龙珠/
	// 如果没有其他路由匹配 /book/龙珠 ，它将重定向至 /book/龙珠/
	router.GET("/book/:title/*author", func(c *gin.Context) {
		title := c.Param("title")
		author := c.Param("author")

		c.String(http.StatusOK, "title: %s, author: %s", title, author)
	})
	_ = router.Run(":8080")
}

func getJsonParams() {
	// router := gin.Default()

}
