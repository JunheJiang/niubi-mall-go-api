package client

import "niubi-mall/service"

type ApiGroup struct {
	IndexApi
	GoodsCategoryApi
	GoodsInfoApi
	ShopCartApi
	UserApi
	UserAddressApi
	OrderApi
}

var mallCarouselService = service.ServiceGroupApp.ClientServiceGroup.CarouselService
var mallGoodsCategoryService = service.ServiceGroupApp.ClientServiceGroup.GoodsCategoryService
var mallUserTokenService = service.ServiceGroupApp.ClientServiceGroup.UserTokenService

var mallGoodsInfoService = service.ServiceGroupApp.ClientServiceGroup.GoodsInfoService
var mallIndexConfigService = service.ServiceGroupApp.ClientServiceGroup.IndexInfoService
var mallUserService = service.ServiceGroupApp.ClientServiceGroup.UserService

var mallUserAddressService = service.ServiceGroupApp.ClientServiceGroup.UserAddressService
var mallShopCartService = service.ServiceGroupApp.ClientServiceGroup.ShopCartService
var mallOrderService = service.ServiceGroupApp.ClientServiceGroup.OrderService
