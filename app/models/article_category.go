package models

type ArticleCategory struct {
	Id         int `json:"id" gorm:"primaryKey;autoIncrement"`
	ArticleId  int `json:"articleId" gorm:"not null;"`
	CategoryId int `json:"categoryId" gorm:"not null;"`
}

func (*ArticleCategory) TableName() string {
	return "tb_article_category"
}
