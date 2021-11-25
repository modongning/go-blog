package main

import (
	"com.otoomo.morningblog-go/app/router"
	sdk "com.otoomo.morningblog-go/sdk"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	gormutil "github.com/modongning/gocommon/gorm"
	"github.com/modongning/gocommon/logger"
	"gorm.io/gorm"
)

func main() {
	logger.InitLogger("./goblog.log")

	r := gin.New()
	mysqlConnectionPool := gormutil.GetGormInstance(logger.GetLogInterface(), "root:123456@tcp(127.0.0.1:13306)/morning_blog?charset=utf8&parseTime=True&loc=Local")
	db := mysqlConnectionPool.Db

	//TODO 初始化gin自定义错误，链路追踪，数据库等

	sdk.Application.SetDB(db.(*gorm.DB))
	sdk.Application.SetEngine(r)

	router.InitRouter()

	if err := r.Run(":8999"); err != nil {
		fmt.Println("启动失败:%d", err.Error())
	}
}
