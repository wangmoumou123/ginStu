package middlewire

import (
	"fmt"
	"ginStu/tools"
	"github.com/gin-gonic/gin"
	"time"
)

func PTime(c *gin.Context) {
	fmt.Println("我是一个中间件")
	start := time.Now().UnixMilli()
	c.Next()
	stop := time.Now().UnixMilli()
	tools.Log("g", "start: %d, stop: %d,use time :%d\n", start, stop, stop-start)
	tools.Log("r", "", start, stop, stop-start)
}
