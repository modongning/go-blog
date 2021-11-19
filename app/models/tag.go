package models

type Tag struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null;comment:标签名称"`
}

func (*Tag) TableName() string{
	return "tb_tag"
}