package manage

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
	"niubi-mall/middleware"
)

type AdminGoodsCategoryRouter struct {
}

func (r *AdminGoodsCategoryRouter) InitManageGoodsCategoryRouter(Router *gin.RouterGroup) {

	goodsCategoryRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())

	var goodsCategoryApi = v1.ApiGroupApp.ManageApiGroup.AdminGoodsCategoryApi
	{
		goodsCategoryRouter.POST("categories", goodsCategoryApi.CreateCategory)
		goodsCategoryRouter.PUT("categories", goodsCategoryApi.UpdateCategory)
		goodsCategoryRouter.GET("categories", goodsCategoryApi.GetCategoryList)
		goodsCategoryRouter.GET("categories/:id", goodsCategoryApi.GetCategory)
		goodsCategoryRouter.DELETE("categories", goodsCategoryApi.DelCategory)
		goodsCategoryRouter.GET("categories4Select", goodsCategoryApi.ListForSelect)
	}

}
