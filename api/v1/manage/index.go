package manage

import "niubi-mall/service"

type Group struct {
	AdminUserApi
}

var mallAdminUserService = service.ServiceGroupApp.ManageServiceGroup.AdminUserService
