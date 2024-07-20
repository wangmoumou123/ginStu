package admin

import (
	"ginStu/modules"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
}

func (con ArticleController) Get(c *gin.Context) {
	var articleList []modules.Article
	modules.DB.Preload("ArticleCate").Find(&articleList)
	//tx := modules.DB.Find(&articleList)
	//fmt.Println(tx.Error)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"content": articleList,
	})

}
