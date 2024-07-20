package admin

import (
	"ginStu/modules"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleCateController struct {
}

func (con ArticleCateController) Get(c *gin.Context) {
	var articleCateList []modules.ArticleCate
	modules.DB.Preload("ArticleList").Find(&articleCateList)
	//tx := modules.DB.Find(&articleList)
	//fmt.Println(tx.Error)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"content": articleCateList,
	})

}
