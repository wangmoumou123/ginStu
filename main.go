package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g.GET("/tt", func(c *gin.Context) {
		c.String(http.StatusOK, "你是: %v 大 大4554声道", "爱惜")
	})
	g.POST("/post_test", func(c *gin.Context) {
		p := c.Query("p")

		fmt.Println(p)
		c.JSON(200, gin.H{
			"status": "ok",
		})

	})
	err := g.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
