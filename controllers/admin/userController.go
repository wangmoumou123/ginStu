package admin

import (
	"fmt"
	"ginStu/modules"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (con UserController) Get(c *gin.Context) {
	userList := []modules.User{}
	//modules.DB.Where("age>12").Find(&userList)
	modules.DB.Where("age>?", 10).Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})

}

func (con UserController) Add(c *gin.Context) {
	u := modules.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusForbidden,
			"err":  err.Error(),
		})
		fmt.Println(err)
		return
	}
	u.AddTime = int(modules.GetUnix())

	modules.DB.Create(&u)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "success"})
	//con.Get(c)
}

func (con UserController) Edit(c *gin.Context) {
	//data, _ := c.GetRawData()
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusForbidden,
			"err":  err.Error(),
		})
	}
	id := params["id"]
	name := params["username"]
	u := modules.User{}
	//1
	modules.DB.Where("id= ?", id).Find(&u).Update("username", name)
	////2
	//modules.DB.Model(&u).Where("id=?", id).Find(&u)
	//u.Username = name.(string)
	//modules.DB.Save(&u)
	////3 先查询后修改保存(忽略)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "success"})

}

func (con UserController) Del2(c *gin.Context) {
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusForbidden,
			"err":  err.Error(),
		})
	}
	name := params["username"]
	u := modules.User{}
	modules.DB.Where("username=?", name).Delete(&u)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "success",
	})

}

func (con UserController) Del(c *gin.Context) {
	defer func() {
		if re := recover(); re != any(nil) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusForbidden,
				"err":  "执行失败",
			})
		}
	}()
	var params map[string]interface{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		return
	}
	name := params["username"]
	if name == nil {
		panic(any("name is nil"))
	}
	u := modules.User{}
	modules.DB.Where("username=?", name).Delete(&u)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"err":  "success",
	})

}
