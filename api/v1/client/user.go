package client

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	request "niubi-mall/model/client/req_param"
	"niubi-mall/model/common/resp_vo"
	"niubi-mall/utils/check"
)

type UserApi struct {
}

func (m *UserApi) UserRegister(c *gin.Context) {
	var req request.RegisterUserParam
	_ = c.ShouldBindJSON(&req)

	if err := check.Verify(req, check.MallUserRegisterVerify); err != nil {
		resp_vo.FailWithMsg(err.Error(), c)
		return
	}
	if err := mallUserService.RegisterUser(req); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		resp_vo.FailWithMsg("创建失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("创建成功", c)
}

func (m *UserApi) UserInfoUpdate(c *gin.Context) {
	var req request.UpdateUserInfoParam
	token := c.GetHeader("token")

	if err := mallUserService.UpdateUserInfo(token, req); err != nil {
		global.GVA_LOG.Error("更新用户信息失败", zap.Error(err))
		resp_vo.FailWithMsg("更新用户信息失败"+err.Error(), c)
	}
	resp_vo.OkWithMsg("更新成功", c)
}

func (m *UserApi) GetUserInfo(c *gin.Context) {
	token := c.GetHeader("token")

	if err, userDetail := mallUserService.GetUserDetail(token); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		resp_vo.FailWithMsg("未查询到记录", c)
	} else {
		resp_vo.OkWithData(userDetail, c)
	}
}

func (m *UserApi) UserLogin(c *gin.Context) {
	var req request.UserLoginParam
	_ = c.ShouldBindJSON(&req)

	if err, _, adminToken := mallUserService.UserLogin(req); err != nil {
		resp_vo.FailWithMsg("登陆失败", c)
	} else {
		resp_vo.OkWithData(adminToken.Token, c)
	}
}

func (m *UserApi) UserLogout(c *gin.Context) {
	token := c.GetHeader("token")

	if err := mallUserTokenService.DeleteMallUserToken(token); err != nil {
		resp_vo.FailWithMsg("登出失败", c)
	} else {
		resp_vo.OkWithMsg("登出成功", c)
	}
}
