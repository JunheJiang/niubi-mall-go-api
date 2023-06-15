package request

import (
	"niubi-mall/model/common/request"
	"niubi-mall/model/manage"
)

type MallUserSearch struct {
	manage.MallUser
	request.PageInfo
}
