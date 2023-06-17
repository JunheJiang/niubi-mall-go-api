package client

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
)

type CarouselIndexRouter struct {
}

func (m *CarouselIndexRouter) InitMallCarouselIndexRouter(Router *gin.RouterGroup) {
	mallCarouselRouter := Router.Group("v1")
	var mallCarouselApi = v1.ApiGroupApp.ClientApiGroup.IndexApi
	{
		mallCarouselRouter.GET("index-infos", mallCarouselApi.MallIndexInfo) // 获取首页数据
	}
}
