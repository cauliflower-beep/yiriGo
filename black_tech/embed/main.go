package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates
var FS embed.FS

//go:embed hello.txt
var s string

func main() {
	fmt.Println(s) // <----- hello.txt 的内容！

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	templ := template.Must(template.New("").ParseFS(FS, "templates/*.html"))
	r.SetHTMLTemplate(templ)

	fe, _ := fs.Sub(FS, "img")
	r.StaticFS("/img", http.FS(fe))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run(":8080")
}
