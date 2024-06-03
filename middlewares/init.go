package middlewares

import (
	"ginStu/tools"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	uName, _ := c.Get("username")
	v, ok := uName.(string)
	if ok {
		tools.Log("g", v)

	} else {
		tools.Log("r", "error")
	}
	c.Next()
	c.Set("age", 18)

}
