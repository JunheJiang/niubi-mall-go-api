package router

import "niubi-mall/router/manage"

type MallGroup struct {
	Admin manage.AdminRouterGroup
}

var RouterGroupApp = new(MallGroup)
