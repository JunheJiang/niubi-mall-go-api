package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"niubi-mall/global"
	"niubi-mall/middleware"
	"niubi-mall/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址

	//Router.Use(middleware.LoadTls())  // 打开就能玩https了

	global.GVA_LOG.Info("use middleware logger")

	// 跨域// 如需跨域可以打开
	Router.Use(middleware.Cors())

	global.GVA_LOG.Info("use middleware cors")

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	// 方便统一添加路由组前缀 多服务器上线使用
	//商城后管路由
	adminRouter := router.RouterGroupApp.Admin
	// admin-api/v1
	AdminGroup := Router.Group("manage-api")
	{
		//商城后管路由初始化
		//用户
		adminRouter.InitManageAdminUserRouter(AdminGroup)
		//轮播图
		adminRouter.InitManageCarouselRouter(AdminGroup)
		//商品分类
		adminRouter.InitManageGoodsCategoryRouter(AdminGroup)
		//商品详情
		adminRouter.InitManageGoodsInfoRouter(AdminGroup)
		//首页配置
		adminRouter.InitManageIndexConfigRouter(AdminGroup)
		//订单
		adminRouter.InitManageOrderRouter(AdminGroup)
		//商品属性

		//物流

		//积分

	}

	//client-api/v1
	//client app api
	clientRouter := router.RouterGroupApp.Client
	ClientGroup := Router.Group("api")
	{
		//client端
		//商城分类
		clientRouter.InitMallGoodsCategoryIndexRouter(ClientGroup)
		//轮播图
		clientRouter.InitMallCarouselIndexRouter(ClientGroup)
		//用户
		clientRouter.InitMallUserRouter(ClientGroup)
		//用户地址
		clientRouter.InitMallUserAddressRouter(ClientGroup)
		//订单
		clientRouter.InitMallOrderRouter(ClientGroup)
		//购物车
		clientRouter.InitMallShopCartRouter(ClientGroup)
		//商品
		clientRouter.InitMallGoodsInfoIndexRouter(ClientGroup)
	}
	return Router
}
