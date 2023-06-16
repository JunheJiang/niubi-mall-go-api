package middleware

import (
	"github.com/gin-gonic/gin"
	"niubi-mall/model/common/resp_vo"
	"niubi-mall/service"
	"time"
)

var manageAdminUserTokenService = service.ServiceGroupApp.AdminServiceGroup.AdminUserTokenService
var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.UserTokenService

func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			resp_vo.FailWithDetail(nil, "未登录或非法访问", c)
			c.Abort()
			return
		}
		err, mallAdminUserToken := manageAdminUserTokenService.ExistAdminToken(token)
		if err != nil {
			resp_vo.FailWithDetail(nil, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if time.Now().After(mallAdminUserToken.ExpireTime) {
			resp_vo.FailWithDetail(nil, "授权已过期", c)
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
			resp_vo.UnLogin(nil, c)
			c.Abort()
			return
		}
		err, mallUserToken := mallUserTokenService.ExistUserToken(token)

		if err != nil {
			resp_vo.UnLogin(nil, c)
			c.Abort()
			return
		}

		if time.Now().After(mallUserToken.ExpireTime) {
			resp_vo.FailWithDetail(nil, "授权已过期", c)
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
