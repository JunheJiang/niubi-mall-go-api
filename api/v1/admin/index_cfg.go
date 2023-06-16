package admin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	manageReq "niubi-mall/model/admin/req_param"
	"niubi-mall/model/common/req_param"
	"niubi-mall/model/common/resp_vo"
	"strconv"
)

type IndexConfigApi struct {
}

func (m *IndexConfigApi) CreateIndexConfig(c *gin.Context) {
	var req manageReq.MallIndexConfigAddParams
	_ = c.ShouldBindJSON(&req)

	if err := mallAdminIndexConfigService.CreateMallIndexConfig(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		resp_vo.FailWithMsg("创建失败"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("创建成功", c)
	}
}

func (m *IndexConfigApi) DeleteIndexConfig(c *gin.Context) {
	var ids req_param.IdsReq
	_ = c.ShouldBindJSON(&ids)

	if err := mallAdminIndexConfigService.DeleteMallIndexConfig(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		resp_vo.FailWithMsg("删除失败"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("删除成功", c)
	}
}

func (m *IndexConfigApi) UpdateIndexConfig(c *gin.Context) {
	var req manageReq.MallIndexConfigUpdateParams
	_ = c.ShouldBindJSON(&req)

	if err := mallAdminIndexConfigService.UpdateMallIndexConfig(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败:"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}
}

func (m *IndexConfigApi) FindIndexConfig(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err, mallIndexConfig := mallAdminIndexConfigService.GetMallIndexConfig(uint(id)); err != nil {
		global.GVA_LOG.Error("查询失败!"+err.Error(), zap.Error(err))
		resp_vo.FailWithMsg("查询失败", c)
	} else {
		resp_vo.OkWithData(mallIndexConfig, c)
	}
}

func (m *IndexConfigApi) GetIndexConfigList(c *gin.Context) {
	var pageInfo manageReq.MallIndexConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallAdminIndexConfigService.GetMallIndexConfigInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
		resp_vo.FailWithMsg("获取失败", c)
	} else {
		resp_vo.OkWithDetail(resp_vo.PageResult{
			List:       list,
			TotalCount: total,
			CurPage:    pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
