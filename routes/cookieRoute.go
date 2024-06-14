package routes

import (
	"ginStu/controllers/sessioncookie"
	"github.com/gin-gonic/gin"
)

func CookieInit(r *gin.Engine) {
	cookieRoute := r.Group("/cookie")
	{
		cookieRoute.GET("/set_c", sessioncookie.Controller{}.SetCookie)
		cookieRoute.GET("/get_c", sessioncookie.Controller{}.GetCookie)
		cookieRoute.GET("/set", sessioncookie.Controller{}.SetSession)
		cookieRoute.GET("/get", sessioncookie.Controller{}.GetSession)

	}
}
