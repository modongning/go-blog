package models

type ArticleTag struct {
	Id        int `json:"id" gorm:"primaryKey;autoIncrement"`
	ArticleId int `json:"articleId" gorm:"not null;"`
	TagId     int `json:"tagId" gorm:"not null;"`
}

func (*ArticleTag) TableName() string {
	return "tb_article_tag"
}
