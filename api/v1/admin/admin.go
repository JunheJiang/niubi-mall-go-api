package admin

import "niubi-mall/service"

type ApiGroup struct {
	UserApi
	CarouselApi
	GoodsCategoryApi
	GoodsInfoApi
	IndexConfigApi
	OrderApi
}

var mallAdminUserService = service.ServiceGroupApp.AdminServiceGroup.AdminUserService
var mallUserService = service.ServiceGroupApp.AdminServiceGroup.UserService
var mallAdminUserTokenService = service.ServiceGroupApp.AdminServiceGroup.AdminUserTokenService
var mallAdminCarouselService = service.ServiceGroupApp.AdminServiceGroup.AdminCarouselService
var mallAdminGoodsCategoryService = service.ServiceGroupApp.AdminServiceGroup.AdminGoodsCategoryService
var mallAdminGoodsInfoService = service.ServiceGroupApp.AdminServiceGroup.AdminGoodsInfoService
var mallAdminIndexConfigService = service.ServiceGroupApp.AdminServiceGroup.AdminIndexConfigService
var mallAdminOrderService = service.ServiceGroupApp.AdminServiceGroup.AdminOrderService
var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
