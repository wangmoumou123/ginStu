package routes

import (
	"ginStu/controllers/index"
	"github.com/gin-gonic/gin"
)

func IndexInit(r *gin.Engine) {
	indexRoute := r.Group("/index")
	{
		indexRoute.GET("/", index.Controller{}.Index)
		indexRoute.GET("/add", index.Controller{}.Add)
		indexRoute.GET("/edit", index.Controller{}.Edit)
	}
}
