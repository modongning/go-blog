package router

import (
	"com.otoomo.morningblog-go/sdk"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := sdk.Application.GetEngine()
	//强制转换类型
	RegisterRouter(engine.(*gin.Engine))
}
