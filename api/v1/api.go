package v1

import "niubi-mall/api/v1/admin"

type ApiGroup struct {
	AdminApiGroup admin.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
