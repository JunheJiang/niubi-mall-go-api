package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	"niubi-mall/model/common/response"
	"niubi-mall/model/manage"
	manageReq "niubi-mall/model/manage/request"
	"niubi-mall/utils/check"
	"niubi-mall/utils/data"
)

type AdminUserApi struct {
}

// CreateAdminUser --- 创建AdminUser
func (m *AdminUserApi) CreateAdminUser(c *gin.Context) {
	var params manageReq.MallAdminParam
	_ = c.ShouldBindJSON(&params)

	if err := check.Verify(params, check.AdminUserRegisterVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	mallAdminUser := manage.MallAdminUser{
		LoginUserName: params.LoginUserName,
		NickName:      params.NickName,
		LoginPassword: data.MD5V([]byte(params.LoginPassword)),
	}

	if err := mallAdminUserService.CreateMallAdminUser(mallAdminUser); err != nil {
		global.GVA_LOG.Error("创建失败:", zap.Error(err))
		response.FailWithMsg("创建失败"+err.Error(), c)
	} else {
		response.OkWithMsg("创建成功", c)
	}
}

// UpdateAdminUserPassword --- 修改密码
func (m *AdminUserApi) UpdateAdminUserPassword(c *gin.Context) {
	var req manageReq.MallUpdatePasswordParam
	_ = c.ShouldBindJSON(&req)
	//mallAdminUserName := manage.MallAdminUser{
	//	LoginPassword: utils.MD5V([]byte(req.LoginPassword)),
	//}
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateMallAdminPassWord(userToken, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMsg("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMsg("更新成功", c)
	}

}

// UpdateAdminUserName --- 更新用户名
func (m *AdminUserApi) UpdateAdminUserName(c *gin.Context) {
	var req manageReq.MallUpdateNameParam
	_ = c.ShouldBindJSON(&req)
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateMallAdminName(userToken, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMsg("更新失败", c)
	} else {
		response.OkWithMsg("更新成功", c)
	}
}

// AdminUserProfile 用id查询AdminUser
func (m *AdminUserApi) AdminUserProfile(c *gin.Context) {
	adminToken := c.GetHeader("token")
	if err, mallAdminUser := mallAdminUserService.GetMallAdminUser(adminToken); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMsg("未查询到记录", c)
	} else {
		mallAdminUser.LoginPassword = "******"
		response.OkWithData(mallAdminUser, c)
	}
}

// AdminLogin 管理员登陆
func (m *AdminUserApi) AdminLogin(c *gin.Context) {
	var adminLoginParams manageReq.MallAdminLoginParam
	_ = c.ShouldBindJSON(&adminLoginParams)
	if err, _, adminToken := mallAdminUserService.AdminLogin(adminLoginParams); err != nil {
		response.FailWithMsg("登陆失败", c)
	} else {
		response.OkWithData(adminToken.Token, c)
	}
}
