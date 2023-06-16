package admin

type AdminServiceGroup struct {
	AdminUserService
	AdminUserTokenService
	UserService
	AdminCarouselService
	AdminGoodsCategoryService
	AdminGoodsInfoService
	AdminIndexConfigService
	AdminOrderService
}
