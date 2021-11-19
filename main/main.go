package main

import (
	"com.otoomo.morningblog-go/app/router"
	sdk "com.otoomo.morningblog-go/sdk"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	r := gin.New()

	db,err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:13306)/morning_blog?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			Logger:logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Info,
					Colorful:      true,
				},
			),
		},
	)
	if err != nil {
		fmt.Println("链接数据库错误",err.Error())
		return
	}

	//TODO 初始化gin自定义错误，链路追踪，数据库等

	sdk.Application.SetDB(db)
	sdk.Application.SetEngine(r)

	router.InitRouter()

	if err := r.Run(":8999");err != nil {
		fmt.Println("启动失败:%d",err.Error())
	}
}