package apis

import (
	"com.otoomo.morningblog-go/app/models"
	"com.otoomo.morningblog-go/app/service"
	"com.otoomo.morningblog-go/app/service/dto"
	cdto "com.otoomo.morningblog-go/common/dto"
	"com.otoomo.morningblog-go/global"
	"com.otoomo.morningblog-go/sdk"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Article struct {
}

func (a *Article) PageQuery(c *gin.Context) {
	global.MyLog.Info("PageQuery")
	article := service.Article{}
	article.Orm = sdk.Application.GetDB()

	req := &dto.ArticlePageRequest{
		CategoryName: c.Param("categoryName"),
	}
	pageNum, _ := c.GetQuery("pageNum")
	pageSize, _ := c.GetQuery("pageSize")
	req.PageNum, _ = strconv.Atoi(pageNum)
	req.PageSize, _ = strconv.Atoi(pageSize)
	result := cdto.NewResult(c)

	var articleList, err = article.PageQuery(req)
	if err != nil {
		result.Fail(err.Error())
		return
	}
	result.Success(articleList)
}

func (a *Article) Save(c *gin.Context) {
	articleReq := &models.Article{}
	c.ShouldBindJSON(articleReq)

	json, _ := json.Marshal(articleReq)
	log.Println(string(json))

	article := &service.Article{}
	article.Orm = sdk.Application.GetDB()

	article.Save(articleReq)

	cdto.NewResult(c).Success("")
}

func (a *Article) Update(c *gin.Context) {
	articleReq := &models.Article{}
	c.ShouldBindJSON(articleReq)

	json, _ := json.Marshal(articleReq)
	log.Println(string(json))

	article := &service.Article{}
	article.Orm = sdk.Application.GetDB()

	result := cdto.NewResult(c)
	if err := article.Update(articleReq); err != nil {
		result.Fail(err.Error())
		return
	}
	result.Success("")
}
