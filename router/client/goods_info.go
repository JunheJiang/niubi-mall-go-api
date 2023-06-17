package client

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
)

type GoodsInfoIndexRouter struct {
}

func (m *GoodsInfoIndexRouter) InitMallGoodsInfoIndexRouter(Router *gin.RouterGroup) {
	mallGoodsRouter := Router.Group("v1")
	var mallGoodsInfoApi = v1.ApiGroupApp.ClientApiGroup.GoodsInfoApi
	{
		mallGoodsRouter.GET("/search", mallGoodsInfoApi.GoodsSearch)           // 商品搜索
		mallGoodsRouter.GET("/goods/detail/:id", mallGoodsInfoApi.GoodsDetail) //商品详情
	}
}
