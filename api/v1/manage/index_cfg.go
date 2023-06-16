package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	"niubi-mall/model/common/request"
	"niubi-mall/model/common/response"
	manageReq "niubi-mall/model/manage/request"
	"strconv"
)

type AdminIndexConfigApi struct {
}

func (m *AdminIndexConfigApi) CreateIndexConfig(c *gin.Context) {
	var req manageReq.MallIndexConfigAddParams
	_ = c.ShouldBindJSON(&req)

	if err := mallAAdminIndexConfigService.CreateMallIndexConfig(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMsg("创建失败"+err.Error(), c)
	} else {
		response.OkWithMsg("创建成功", c)
	}
}

func (m *AdminIndexConfigApi) DeleteIndexConfig(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)

	if err := mallAAdminIndexConfigService.DeleteMallIndexConfig(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMsg("删除失败"+err.Error(), c)
	} else {
		response.OkWithMsg("删除成功", c)
	}
}

func (m *AdminIndexConfigApi) UpdateIndexConfig(c *gin.Context) {
	var req manageReq.MallIndexConfigUpdateParams
	_ = c.ShouldBindJSON(&req)

	if err := mallAAdminIndexConfigService.UpdateMallIndexConfig(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMsg("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMsg("更新成功", c)
	}
}

func (m *AdminIndexConfigApi) FindIndexConfig(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err, mallIndexConfig := mallAAdminIndexConfigService.GetMallIndexConfig(uint(id)); err != nil {
		global.GVA_LOG.Error("查询失败!"+err.Error(), zap.Error(err))
		response.FailWithMsg("查询失败", c)
	} else {
		response.OkWithData(mallIndexConfig, c)
	}
}

func (m *AdminIndexConfigApi) GetIndexConfigList(c *gin.Context) {
	var pageInfo manageReq.MallIndexConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallAAdminIndexConfigService.GetMallIndexConfigInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
		response.FailWithMsg("获取失败", c)
	} else {
		response.OkWithDetail(response.PageResult{
			List:       list,
			TotalCount: total,
			CurPage:    pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
