package router

import "niubi-mall/router/manage"

type Group struct {
	Manage manage.AdminUserRouter
}

var RouterGroupApp = new(Group)
