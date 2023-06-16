package middleware

import (
	"github.com/gin-gonic/gin"
	"niubi-mall/model/common/response"
	"niubi-mall/service"
	"time"
)

var manageAdminUserTokenService = service.ServiceGroupApp.AdminServiceGroup.AdminUserTokenService
var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.UserTokenService

func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetail(nil, "未登录或非法访问", c)
			c.Abort()
			return
		}
		err, mallAdminUserToken := manageAdminUserTokenService.ExistAdminToken(token)
		if err != nil {
			response.FailWithDetail(nil, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if time.Now().After(mallAdminUserToken.ExpireTime) {
			response.FailWithDetail(nil, "授权已过期", c)
			err = manageAdminUserTokenService.DeleteMallAdminUserToken(token)
			if err != nil {
				return
			}
			c.Abort()
			return
		}
		c.Next()
	}

}

func UserJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.UnLogin(nil, c)
			c.Abort()
			return
		}
		err, mallUserToken := mallUserTokenService.ExistUserToken(token)

		if err != nil {
			response.UnLogin(nil, c)
			c.Abort()
			return
		}

		if time.Now().After(mallUserToken.ExpireTime) {
			response.FailWithDetail(nil, "授权已过期", c)
			err = mallUserTokenService.DeleteMallUserToken(token)
			if err != nil {
				return
			}
			c.Abort()
			return
		}
		c.Next()
	}

}
