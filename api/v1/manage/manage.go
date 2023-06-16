package manage

import "niubi-mall/service"

type Group struct {
	AdminUserApi
}

var mallAdminUserService = service.ServiceGroupApp.ManageServiceGroup.AdminUserService
var mallUserService = service.ServiceGroupApp.ManageServiceGroup.UserService
var mallAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.AdminUserTokenService
var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
