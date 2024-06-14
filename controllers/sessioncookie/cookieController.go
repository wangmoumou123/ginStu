package sessioncookie

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (con Controller) SetCookie(c *gin.Context) {
	_, err := c.Cookie("c_name")
	if err == nil {
		c.Redirect(http.StatusFound, "get_c")
		return

	}
	// 设置cookie
	c.SetCookie("c_name", "is_login", 10, "/", "*", false, false)

	c.HTML(http.StatusOK, "cookie/index.html", gin.H{
		"status": "set cookie",
	})
	//c.JSON(http.StatusOK, gin.H{
	//	"status": "set cookie",
	//})

}

func (con Controller) GetCookie(c *gin.Context) {
	//获取cookie
	isLogin, err := c.Cookie("c_name")
	if err != nil {
		c.Redirect(http.StatusFound, "set_c")
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "get cookie",
		"is_login": isLogin,
	})

}

func (con Controller) SetSession(c *gin.Context) {
	// 设置session
	session := sessions.Default(c)
	if session.Get("is_login") != nil {
		c.Redirect(http.StatusFound, "get")
		return
	}
	session.Options(sessions.Options{
		MaxAge: 10,
	})
	session.Set("is_login", 1)
	err := session.Save()
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "cookie/index.html", gin.H{
		"status": "set session",
	})
}

func (con Controller) GetSession(c *gin.Context) {
	// 获取session
	session := sessions.Default(c)
	if session.Get("is_login") != nil {
		c.HTML(http.StatusOK, "cookie/index.html", gin.H{
			"status": "已登录",
		})
	} else {
		//c.HTML(http.StatusOK, "cookie/index.html", gin.H{
		//	"status":   "未登录",
		//	"is_login": "http://127.0.0.1:8080/cookie/set",
		//})
		c.Redirect(http.StatusFound, "set")
	}
}
