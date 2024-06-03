package routes

import (
	"ginStu/controllers/action"
	"github.com/gin-gonic/gin"
)

func ActionInit(r *gin.Engine) {
	actionRoute := r.Group("/action")
	{
		actionRoute.GET("/WebGetTtInfo", action.Controller{}.WebGetTtInfo)

	}
}
