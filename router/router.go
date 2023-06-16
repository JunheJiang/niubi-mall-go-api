package router

import "niubi-mall/router/admin"

type MallGroup struct {
	Admin admin.RouterGroup
}

var RouterGroupApp = new(MallGroup)
