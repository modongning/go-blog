package models

import "time"

type CommonModelField struct {
	CreatedBy   int       `json:"createdBy" gorm:"size:20,comment:创建人"`
	CreatedTime time.Time `json:"createdTime" gorm:"default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedBy   int       `json:"createdBy" gorm:"comment:更新人"`
	UpdatedTime time.Time `json:"createdTime" gorm:"default:CURRENT_TIMESTAMP;comment:更新时间"`
}

type Status struct {
	Status int8 `json:"status" gorm:"default:1;comment:状态,0-无效,1-正常/有效"`
}

type DelStatus struct {
	DelStatus int8 `json:"status" gorm:"default:0;comment:删除状态,0-未删除,1-已删除"`
}
