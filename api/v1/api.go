package v1

import (
	"niubi-mall/api/v1/admin"
	"niubi-mall/api/v1/client"
)

type ApiGroup struct {
	AdminApiGroup  admin.ApiGroup
	ClientApiGroup client.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
