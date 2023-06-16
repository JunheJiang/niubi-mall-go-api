package client

import (
	"github.com/gin-gonic/gin"
	v1 "niubi-mall/api/v1"
	"niubi-mall/middleware"
)

type UserRouter struct {
}

func (m *UserRouter) InitMallUserRouter(Router *gin.RouterGroup) {

	mallUserRouter := Router.Group("v1").Use(middleware.UserJWTAuth())

	userRouter := Router.Group("v1")

	var mallUserApi = v1.ApiGroupApp.ClientApiGroup.UserApi
	{
		mallUserRouter.PUT("/user/info", mallUserApi.UserInfoUpdate) //修改用户信息
		mallUserRouter.GET("/user/info", mallUserApi.GetUserInfo)    //获取用户信息
		mallUserRouter.POST("/user/logout", mallUserApi.UserLogout)  //登出
	}
	{
		userRouter.POST("/user/register", mallUserApi.UserRegister) //用户注册
		userRouter.POST("/user/login", mallUserApi.UserLogin)       //登陆

	}

}
