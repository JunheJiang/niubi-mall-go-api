package client

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
)

type GoodsCategoryIndexRouter struct {
}

func (m *GoodsCategoryIndexRouter) InitMallGoodsCategoryIndexRouter(Router *gin.RouterGroup) {
	mallGoodsRouter := Router.Group("v1")
	var mallGoodsCategoryApi = v1.ApiGroupApp.ClientApiGroup.GoodsCategoryApi
	{
		mallGoodsRouter.GET("categories", mallGoodsCategoryApi.GetGoodsCategory) // 获取分类数据
	}
}
