package req_param

import (
	"niubi-mall/model/admin/db_entity"
	"niubi-mall/model/common/req_param"
)

type MallUserSearch struct {
	db_entity.MallUser
	req_param.PageInfo
}
