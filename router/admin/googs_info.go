package admin

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
	"niubi-mall/middleware"
)

type GoodsInfoRouter struct {
}

func (m *GoodsInfoRouter) InitManageGoodsInfoRouter(Router *gin.RouterGroup) {

	mallGoodsInfoRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())

	var mallGoodsInfoApi = v1.ApiGroupApp.AdminApiGroup.AdminGoodsInfoApi

	{
		mallGoodsInfoRouter.POST("goods", mallGoodsInfoApi.CreateGoodsInfo)                    // 新建MallGoodsInfo
		mallGoodsInfoRouter.DELETE("deleteMallGoodsInfo", mallGoodsInfoApi.DeleteGoodsInfo)    // 删除MallGoodsInfo
		mallGoodsInfoRouter.PUT("goods/status/:status", mallGoodsInfoApi.ChangeGoodsInfoByIds) // 上下架
		mallGoodsInfoRouter.PUT("goods", mallGoodsInfoApi.UpdateGoodsInfo)                     // 更新MallGoodsInfo
		mallGoodsInfoRouter.GET("goods/:id", mallGoodsInfoApi.FindGoodsInfo)                   // 根据ID获取MallGoodsInfo
		mallGoodsInfoRouter.GET("goods/list", mallGoodsInfoApi.GetGoodsInfoList)               // 获取MallGoodsInfo列表
	}
}
