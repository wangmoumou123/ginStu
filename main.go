package main

import (
	"ginStu/routes"
	"ginStu/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store1 := cookie.NewStore([]byte("wst123456"))

	store, _ := redis.NewStore(1000, "tcp", "127.0.0.1:6379", "", []byte("qrwegfdscdvd55"))
	tools.Log("g", store1)
	r.Use(sessions.Sessions("mySession", store))
	r.LoadHTMLGlob("templates/*/**")
	r.MaxMultipartMemory = 8 << 20
	routes.AdminInit(r)
	routes.ArticleInit(r)
	routes.IndexInit(r)
	routes.ActionInit(r)
	routes.CookieInit(r)
	err := r.Run()
	if err != nil {

		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
