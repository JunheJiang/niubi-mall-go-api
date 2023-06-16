package req_param

import (
	"niubi-mall/model/common/req_param"
	"niubi-mall/model/manage/db_entity"
)

type MallUserSearch struct {
	db_entity.MallUser
	req_param.PageInfo
}
