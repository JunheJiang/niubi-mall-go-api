package admin

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
	"niubi-mall/middleware"
)

type UserRouter struct {
}

func (r *UserRouter) InitManageAdminUserRouter(Router *gin.RouterGroup) {
	mallAdminUserRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())

	mallAdminUserWithoutRouter := Router.Group("v1")

	var mallAdminUserApi = v1.ApiGroupApp.AdminApiGroup.AdminUserApi
	{
		mallAdminUserRouter.POST("createMallAdminUser", mallAdminUserApi.CreateAdminUser) // 新建MallAdminUser
		mallAdminUserRouter.PUT("adminUser/name", mallAdminUserApi.UpdateAdminUserName)   // 更新MallAdminUser
		mallAdminUserRouter.PUT("adminUser/password", mallAdminUserApi.UpdateAdminUserPassword)
	}
	{
		mallAdminUserWithoutRouter.POST("adminUser/login", mallAdminUserApi.AdminLogin) //管理员登陆
	}
}
