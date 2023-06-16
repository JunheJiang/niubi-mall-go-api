package admin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	"niubi-mall/model/admin/db_entity"
	manageReq "niubi-mall/model/admin/req_param"
	"niubi-mall/model/common/req_param"
	"niubi-mall/model/common/resp_vo"
	"niubi-mall/model/example"
	"niubi-mall/utils/check"
	"niubi-mall/utils/data"
	"strconv"
)

type AdminUserApi struct {
}

// CreateAdminUser --- 创建AdminUser
func (m *AdminUserApi) CreateAdminUser(c *gin.Context) {
	var params manageReq.MallAdminParam
	_ = c.ShouldBindJSON(&params)

	if err := check.Verify(params, check.AdminUserRegisterVerify); err != nil {
		resp_vo.FailWithMsg(err.Error(), c)
		return
	}

	mallAdminUser := db_entity.MallAdminUser{
		LoginUserName: params.LoginUserName,
		NickName:      params.NickName,
		LoginPassword: data.MD5V([]byte(params.LoginPassword)),
	}

	if err := mallAdminUserService.CreateMallAdminUser(mallAdminUser); err != nil {
		global.GVA_LOG.Error("创建失败:", zap.Error(err))
		resp_vo.FailWithMsg("创建失败"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("创建成功", c)
	}
}

// UpdateAdminUserPassword --- 修改密码
func (m *AdminUserApi) UpdateAdminUserPassword(c *gin.Context) {
	var req manageReq.MallUpdatePasswordParam
	_ = c.ShouldBindJSON(&req)
	//mallAdminUserName := admin.MallAdminUser{
	//	LoginPassword: utils.MD5V([]byte(req.LoginPassword)),
	//}
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateMallAdminPassWord(userToken, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败:"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}

}

// UpdateAdminUserName --- 更新用户名
func (m *AdminUserApi) UpdateAdminUserName(c *gin.Context) {
	var req manageReq.MallUpdateNameParam
	_ = c.ShouldBindJSON(&req)
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateMallAdminName(userToken, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败", c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}
}

// AdminUserProfile 用id查询AdminUser
func (m *AdminUserApi) AdminUserProfile(c *gin.Context) {
	adminToken := c.GetHeader("token")
	if err, mallAdminUser := mallAdminUserService.GetMallAdminUser(adminToken); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		resp_vo.FailWithMsg("未查询到记录", c)
	} else {
		mallAdminUser.LoginPassword = "******"
		resp_vo.OkWithData(mallAdminUser, c)
	}
}

// AdminLogin 管理员登陆
func (m *AdminUserApi) AdminLogin(c *gin.Context) {
	var adminLoginParams manageReq.MallAdminLoginParam
	_ = c.ShouldBindJSON(&adminLoginParams)
	if err, _, adminToken := mallAdminUserService.AdminLogin(adminLoginParams); err != nil {
		resp_vo.FailWithMsg("登陆失败", c)
	} else {
		resp_vo.OkWithData(adminToken.Token, c)
	}
}

// UserList --- 注册用户列表
func (m *AdminUserApi) UserList(c *gin.Context) {
	var pageInfo manageReq.MallUserSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := mallUserService.GetMallUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		resp_vo.FailWithMsg("获取失败!", c)
	} else {
		resp_vo.OkWithDetail(resp_vo.PageResult{
			List:       list,
			TotalCount: total,
			CurPage:    pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (m *AdminUserApi) LockUser(c *gin.Context) {
	lockStatus, _ := strconv.Atoi(c.Param("lockStatus"))
	var IDS req_param.IdsReq
	_ = c.ShouldBindJSON(&IDS)

	if err := mallUserService.LockUser(IDS, lockStatus); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败!", c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}
}

// UploadFile 上传单图
// 此处上传图片的功能可用，但是原前端项目的图片链接为服务器地址，如需要显示图片，需要修改前端指向的图片链接
func (m *AdminUserApi) UploadFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		resp_vo.FailWithMsg("接收文件失败", c)
		return
	}
	err, file = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		resp_vo.FailWithMsg("修改数据库链接失败", c)
		return
	}
	//这里直接使用本地的url
	resp_vo.OkWithData("http://localhost:8888/"+file.Url, c)
}
