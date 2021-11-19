package service

import (
	"com.otoomo.morningblog-go/app/models"
	"com.otoomo.morningblog-go/app/service/dto"
	"com.otoomo.morningblog-go/sdk/service"
	"errors"
)

type Article struct {
	service.Service
}

/*
PageQuery 分页查询
 */
func (a *Article) PageQuery(articlePageRes *dto.ArticlePageRequest) (articleList []models.Article, err error) {
	articleList = make([]models.Article, 0)
	db := a.Orm.Table("tb_article a")
	db.Select("a.*")
	db.Joins("left join tb_article_category ac on a.category = ac.id")

	categoryName := articlePageRes.CategoryName
	if categoryName != "" && len(categoryName) > 0 {
		db.Where("ac.category id = '%?%'", articlePageRes.CategoryName)
	}
	offset := (articlePageRes.PageNum - 1) * articlePageRes.PageSize
	db.Limit(articlePageRes.PageSize).Offset(offset)

	if err := db.Find(&articleList).Error; err != nil {
		return nil, err
	}
	return
}

/*
Save 添加
 */
func (a *Article) Save(article *models.Article) {
	a.Orm.Save(article)
}

/*
Update 更新
 */
func (a *Article) Update(req *models.Article) error {
	if id := req.Id; id == 0 {
		return errors.New("参数错误")
	}
	a.Orm.Updates(req)
	return nil
}
