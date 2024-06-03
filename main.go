package main

import (
	"ginStu/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.AdminInit(r)
	routes.ArticleInit(r)
	routes.IndexInit(r)
	routes.ActionInit(r)
	err := r.Run()
	if err != nil {

		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
