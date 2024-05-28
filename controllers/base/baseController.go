package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string      `json:"title"`
	Content string      `json:"content"`
	Alis    interface{} `json:"alis"`
}
type Tt struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name"`
}

type Controller struct {
}

func (con Controller) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key":   "base",
		"value": "index",
	})
}

func (con Controller) Add(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key":   "base",
		"value": "add",
	})
}
func (con Controller) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key":   "base",
		"value": "edit",
	})
}

func (con Controller) TT(c *gin.Context) {
	c.String(http.StatusOK, "你是: %v 大 大4554声道", "爱惜")

}
func (con Controller) PostTest(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		return
	}
	//Log(data, "r")
	fmt.Printf("%s===%T\n", data, data)
	c.JSON(200, gin.H{
		"status": "ok",
	})

}

func (con Controller) Post(c *gin.Context) {

	var tt Tt
	if err := c.ShouldBindJSON(&tt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
		})
		fmt.Println(err)
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"tt":     tt,
	})
}

//
//func (con Controller) Index(c *gin.Context) {
//	art := Article{"123", "456789", map[string]string{"1": "一", "2": "二", "3": "三"}}
//	c.HTML(http.StatusOK, "default/index.html", gin.H{
//		"title": "hello",
//		"art":   art,
//	})
//}

func (con Controller) AIndex(c *gin.Context) {
	art := Article{"123", "456789", []string{"1", "2", "3"}}
	c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{
		"title": "hello",
		"art":   art,
	})
}
func (con Controller) Middle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "middle",
	})
}
