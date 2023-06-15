package v1

import "niubi-mall/api/v1/manage"

type ApiGroup struct {
	ManageApiGroup manage.Group
}

var ApiGroupApp = new(ApiGroup)
