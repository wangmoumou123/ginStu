package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/tt", func(c *gin.Context) {
		c.String(http.StatusOK, "你是: %v 大 大4554声道", "爱惜")
	})
	r.POST("/post_test", func(c *gin.Context) {
		p := c.Query("p")

		fmt.Println(p)
		c.JSON(200, gin.H{
			"status": "ok",
		})

	})
	r.GET("/index", func(c *gin.Context) {
		art := Article{"123", "456789"}
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "hello",
			"art":   art,
		})
	})
	r.GET("/a_index", func(c *gin.Context) {
		art := Article{"123", "456789"}
		c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{
			"title": "hello",
			"art":   art,
		})
	})

	err := r.Run()
	if err != nil {

		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
