package models

type Category struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null;comment:分类名称"`
}

func (*Category) TableName() string{
	return "tb_category"
}