package article

import (
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
		"key":   "article",
		"value": "index",
	})
}

func (con Controller) Add(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key":   "article",
		"value": "add",
	})
}
func (con Controller) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key":   "article",
		"value": "edit",
	})
}
