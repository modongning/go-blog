package dto

import "com.otoomo.morningblog-go/common/dto"

type ArticlePageRequest struct {
	CategoryName string `json:"categoryName"`
	dto.PageDTO
}
