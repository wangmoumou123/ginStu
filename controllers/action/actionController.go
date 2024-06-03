package action

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
		"key":   "index",
		"value": "index",
	})
}

func (con Controller) Add(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key":   "index",
		"value": "add",
	})
}
func (con Controller) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key":   "index",
		"value": "edit",
	})
}
func (con Controller) WebGetTtInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"extend":                 "extend",
		"cmdCode":                3001,
		"ttServiceSwitch":        true,
		"ttNetworkState":         "å…¥ç½‘",
		"ttNetworkFrequencyInfo": 381,
		"ttSignalQuality":        -76,
		"ttSnr":                  183,
		"ttSimInfo":              "460590104269269",
		"ttSmsCenterNum":         "+8617400010200",
		"ttSimUnlockStateSwitch": false,
		"ttPsStateSwitch":        true,
		"ttPsState":              "æ¿€æ´»ä¸­",
		"ttIpAddr":               "Fail",
		"ttModel":                1,
		"ttMul":                  384,
		"ttMdl":                  384,
		"ttGul":                  0,
		"ttGdl":                  0,
		"ttSpeedModelType":       2,
		"ttWaveVelocitySwitch":   true,
	})
}
