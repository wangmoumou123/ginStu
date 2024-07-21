package main

import (
	"fmt"
	"ginStu/modules"
	"ginStu/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	mysqlIp := cfg.Section("mysql").Key("ip").String()
	mysqlPort := cfg.Section("mysql").Key("port").String()
	mysqlUser := cfg.Section("mysql").Key("user").String()
	mysqlPwd := cfg.Section("mysql").Key("password").String()
	mysqlDatabase := cfg.Section("mysql").Key("database").String()

	redisIp := cfg.Section("redis").Key("ip").String()
	redisPort := cfg.Section("redis").Key("port").String()
	cfg.Section("").Key("class").SetValue("wst")

	err = cfg.SaveTo("./conf/app.ini")
	if err != nil {
		return
	}
	mysqlDSN := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPwd, mysqlIp, mysqlPort, mysqlDatabase)
	redisDSN := fmt.Sprintf("%v:%v", redisIp, redisPort)

	modules.SetDSN(mysqlDSN, redisDSN)

	r := gin.Default()
	r.Use(sessions.Sessions("mySession", modules.Store))
	r.LoadHTMLGlob("templates/*/**")
	r.MaxMultipartMemory = 8 << 20
	routes.AdminInit(r)
	routes.ArticleInit(r)
	routes.IndexInit(r)
	routes.ActionInit(r)
	routes.CookieInit(r)
	err = r.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
