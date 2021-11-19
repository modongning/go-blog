package router

import "github.com/gin-gonic/gin"

var (
	routerGroup = make([]func(group *gin.RouterGroup), 0)
)

func RegisterRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	for _, f := range routerGroup {
		f(api)
	}
}
