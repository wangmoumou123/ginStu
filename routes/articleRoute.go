package routes

import (
	"ginStu/controllers/article"
	"github.com/gin-gonic/gin"
)

func ArticleInit(r *gin.Engine) {
	adminRoute := r.Group("/article")
	{
		adminRoute.GET("/", article.Controller{}.Index)
		adminRoute.GET("/add", article.Controller{}.Add)
		adminRoute.GET("/edit", article.Controller{}.Edit)
	}
}
