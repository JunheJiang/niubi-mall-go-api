package router

import (
	"niubi-mall/router/admin"
	"niubi-mall/router/client"
)

type MallGroup struct {
	Admin  admin.RouterGroup
	Client client.RouterGroup
}

var RouterGroupApp = new(MallGroup)
