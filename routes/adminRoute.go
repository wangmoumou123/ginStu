package routes

import (
	"ginStu/controllers/admin"
	"ginStu/middlewire"
	"github.com/gin-gonic/gin"
)

func AdminInit(r *gin.Engine) {
	adminRoute := r.Group("/admin")
	{
		adminRoute.GET("/", middlewire.PTime, admin.Controller{}.Index)
		adminRoute.GET("/add", admin.Controller{}.Add)
		adminRoute.GET("/edit", admin.Controller{}.Edit)
	}
}
