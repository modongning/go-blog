package router

import (
	"com.otoomo.morningblog-go/app/apis"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup = append(routerGroup, func(group *gin.RouterGroup) {
		api := apis.Article{}
		r := group.Group("/article")
		{
			r.GET("/pageQuery", api.PageQuery)
			r.POST("", api.Save)
			r.PUT("", api.Update)
		}
	})
}
