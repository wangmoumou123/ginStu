package middlewares

import (
	"ginStu/tools"
	"github.com/gin-gonic/gin"
	"time"
)

func PTime(c *gin.Context) {
	tools.Log("", "我是一个中间件")
	start := time.Now().UnixMilli()
	c.Set("username", "wst")

	c.Next()
	time.Sleep(time.Nanosecond)
	stop := time.Now().UnixMilli()
	tools.Log("g", " "+
		"start: ", time.UnixMilli(start).Format("2006-01-02 03:04:05"),
		" stop: ", time.UnixMilli(stop).Format("2006-01-02 03:04:05"),
		" use time : ", stop-start)
	tools.Log("g", time.Now().Format("2006-01-02 03:04:05"))
	tools.Log("g", "uri==", c.Request.RequestURI)
	age, _ := c.Get("age")
	a, ok := age.(int)
	if ok {
		tools.Log("g", "age===>", a)
	} else {
		tools.Log("r", "error")
	}
}
