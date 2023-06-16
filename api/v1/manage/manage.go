package manage

import "niubi-mall/service"

type ApiGroup struct {
	AdminUserApi
	AdminCarouselApi
	AdminGoodsCategoryApi
}

var mallAdminUserService = service.ServiceGroupApp.AdminServiceGroup.AdminUserService
var mallUserService = service.ServiceGroupApp.AdminServiceGroup.UserService
var mallAdminUserTokenService = service.ServiceGroupApp.AdminServiceGroup.AdminUserTokenService
var mallAdminCarouselService = service.ServiceGroupApp.AdminServiceGroup.AdminCarouselService
var mallAdminGoodsCategoryService = service.ServiceGroupApp.AdminServiceGroup.AdminGoodsCategoryService
var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
