package admin

import (
	"crypto/md5"
	"fmt"
	"ginStu/controllers/base"
	"ginStu/modules"
	"ginStu/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
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
	base.Controller
}

func (con Controller) Index(c *gin.Context) {
	/*go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("go 执行")
		t := time.Now()
		fmt.Println(t, "===>", modules.UntoTime(int(t.Unix())))
	}()*/
	d := []byte("12345")
	dEn := md5.Sum([]byte(d))
	//dDn := md5.New()
	fmt.Printf("d===>%x, den=====>%x\n", d, dEn)
	c.JSON(http.StatusOK, gin.H{
		"key":   "admin",
		"value": "index",
	})
}

func (con Controller) Add(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"key":   "admin",
	//	"value": "add",
	//})
	con.Controller.Add(c)
}
func (con Controller) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key":   "admin",
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

func (con Controller) UploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/upload.html", gin.H{
		"status": "ok",
	})

}

func (con Controller) Upload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		return
	}
	err = c.SaveUploadedFile(f, "file/"+f.Filename)
	if err != nil {
		return
	}

	c.String(http.StatusOK, "ok")
}

func (con Controller) UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		return
	}
	allowMap := map[string]bool{
		".jpg": true,
		".png": true,
		".img": true,
		".wav": true,
	}
	files := form.File["file[]"]
	dst := path.Join("file/", modules.GetDay())
	err = os.MkdirAll(dst, 0666)
	if err != nil {
		tools.Log("r", err)
		return
	}
	tools.Log("g", len(files))
	for _, f := range files {
		extName := path.Ext(f.Filename)
		tools.Log("r", extName)
		if _, ok := allowMap[extName]; !ok {
			c.String(http.StatusOK, "file type invalid!")
			return
		}
		tools.Log("r", 11111)

		tools.Log("r", 22222)

		err = c.SaveUploadedFile(f, path.Join(dst, strconv.FormatInt(modules.GetUnix(), 10)+extName))
		if err != nil {
			return
		}
	}

	//err = c.SaveUploadedFile(f, "file/"+f.Filename)
	//if err != nil {
	//	return
	//}

	c.String(http.StatusOK, "ok")
}
