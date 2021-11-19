package models

import (
	"gorm.io/gorm"
	"time"
)

/*
Article 文章数据库实体
*/
type Article struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null;size:400;comment:标题"`
	Cover       string    `json:"cover" gorm:"size:400;comment:封面图片"`
	Author      string    `json:"author" gorm:"size:100;comment:作者"`
	Content     string    `json:"content" gorm:"type:mediumtext;not null;comment:html内容"`
	ContentMd   string    `json:"contentMd" gorm:"type:mediumtext;not null;comment:markdown内容"`
	Category    string    `json:"category" gorm:"size:100;comment:分类"`
	PublishTime time.Time `json:"publishTime" gorm:"comment:发布时间"`
	State       int       `json:"state"  gorm:"comment:状态"`
	CreateTime  time.Time `json:"createTime" gorm:"default:CURRENT_TIMESTAMP;comment:创建时间"`
	EditTime    time.Time `json:"editTime" gorm:"default:CURRENT_TIMESTAMP;comment:更新时间"`
}

/*
TableName 设置表名
*/
func (*Article) TableName() string {
	return "tb_article"
}

/*
BeforeCreate 保存前初始化时间信息
 */
func (a * Article) BeforeCreate(*gorm.DB) (err error){
	a.CreateTime = time.Now()
	a.EditTime = time.Now()
	return
}

func (a * Article) BeforeUpdate(*gorm.DB) (err error){
	a.EditTime = time.Now()
	return
}