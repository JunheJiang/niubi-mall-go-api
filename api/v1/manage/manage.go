package manage

import "niubi-mall/service"

type Group struct {
	AdminUserApi
	AdminCarouselApi
}

var mallAdminUserService = service.ServiceGroupApp.ManageServiceGroup.AdminUserService
var mallUserService = service.ServiceGroupApp.ManageServiceGroup.UserService
var mallAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.AdminUserTokenService
var mallAdminCarouselService = service.ServiceGroupApp.ManageServiceGroup.AdminCarouselService
var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
