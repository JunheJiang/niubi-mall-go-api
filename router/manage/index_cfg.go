package manage

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
	"niubi-mall/middleware"
)

type AdminIndexConfigRouter struct {
}

func (r *AdminIndexConfigRouter) InitManageIndexConfigRouter(Router *gin.RouterGroup) {

	mallIndexConfigRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())

	var mallIndexConfigApi = v1.ApiGroupApp.ManageApiGroup.AdminIndexConfigApi
	{
		mallIndexConfigRouter.POST("indexConfigs", mallIndexConfigApi.CreateIndexConfig)        // 新建MallIndexConfig
		mallIndexConfigRouter.POST("indexConfigs/delete", mallIndexConfigApi.DeleteIndexConfig) // 删除MallIndexConfig
		mallIndexConfigRouter.PUT("indexConfigs", mallIndexConfigApi.UpdateIndexConfig)         // 更新MallIndexConfig
		mallIndexConfigRouter.GET("indexConfigs/:id", mallIndexConfigApi.FindIndexConfig)       // 根据ID获取MallIndexConfig
		mallIndexConfigRouter.GET("indexConfigs", mallIndexConfigApi.GetIndexConfigList)        // 获取MallIndexConfig列表
	}
}
