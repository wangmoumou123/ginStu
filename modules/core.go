package modules

import (
	"fmt"
	"github.com/gin-contrib/sessions/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var Store redis.Store

func Init(mysqlDSN, redisDSN string) {

	// mysql
	//dsn := "ws:1234@tcp(192.168.0.56:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{
		//SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	//cookie存储
	//store1 := cookie.NewStore([]byte("wst123456"))
	//tools.Log("g", store1)

	//redis存储
	//store, _ := redis.NewStore(1000, "tcp", "127.0.0.1:6379", "", []byte("qrwegfdscdvd55"))
	Store, _ = redis.NewStore(1000, "tcp", redisDSN, "", []byte("qrwegfdscdvd55"))
}

func SetDSN(mysqlDSN, redisDSN string) {
	Init(mysqlDSN, redisDSN)
}
