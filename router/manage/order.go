package manage

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
	"niubi-mall/middleware"
)

type AdminOrderRouter struct {
}

func (r *AdminOrderRouter) InitManageOrderRouter(Router *gin.RouterGroup) {

	mallOrderRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())

	var mallOrderApi = v1.ApiGroupApp.ManageApiGroup.AdminOrderApi
	{
		mallOrderRouter.PUT("orders/checkDone", mallOrderApi.CheckDoneOrder) // 发货
		mallOrderRouter.PUT("orders/checkOut", mallOrderApi.CheckOutOrder)   // 出库
		mallOrderRouter.PUT("orders/close", mallOrderApi.CloseOrder)         // 出库
		mallOrderRouter.GET("orders/:orderId", mallOrderApi.FindMallOrder)   // 根据ID获取MallOrder
		mallOrderRouter.GET("orders", mallOrderApi.GetMallOrderList)         // 获取MallOrder列表
	}
}
