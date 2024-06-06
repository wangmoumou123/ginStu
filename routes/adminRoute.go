package routes

import (
	"ginStu/controllers/admin"
	"ginStu/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminInit(r *gin.Engine) {
	adminRoute := r.Group("/admin")
	{
		adminRoute.Use(middlewares.PTime, middlewares.InitMiddleware)

		adminRoute.GET("/", admin.Controller{}.Index)
		adminRoute.GET("/add", admin.Controller{}.Add)
		adminRoute.GET("/edit", admin.Controller{}.Edit)
		adminRoute.GET("/upload", admin.Controller{}.UploadPage)
		//adminRoute.POST("/upload", admin.Controller{}.Upload)
		adminRoute.POST("/upload", admin.Controller{}.UploadFiles)
	}
	//time.Now().UnixNano()
}
