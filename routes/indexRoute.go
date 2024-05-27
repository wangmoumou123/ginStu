package routes

import (
	"ginStu/controllers/index"
	"github.com/gin-gonic/gin"
)

func IndexInit(r *gin.Engine) {
	adminRoute := r.Group("/index")
	{
		adminRoute.GET("/", index.Controller{}.Index)
		adminRoute.GET("/add", index.Controller{}.Add)
		adminRoute.GET("/edit", index.Controller{}.Edit)
	}
}
