package req_param

import (
	"niubi-mall/model/common"
	"niubi-mall/model/common/req_param"
)

type MallGoodsCategoryReq struct {
	CategoryId    int             `json:"categoryId"`
	CategoryLevel int             `json:"categoryLevel" `
	ParentId      int             `json:"parentId"`
	CategoryName  string          `json:"categoryName" `
	CategoryRank  string          `json:"categoryRank" `
	IsDeleted     int             `json:"isDeleted" `
	CreateTime    common.JSONTime `json:"createTime" ` // 创建时间
	UpdateTime    common.JSONTime `json:"updateTime" ` // 更新时间
}

type SearchCategoryParams struct {
	CategoryLevel int `json:"categoryLevel" form:"categoryLevel"`
	ParentId      int `json:"parentId" form:"parentId"`
	req_param.PageInfo
}
