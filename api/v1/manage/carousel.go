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

type AdminCarouselApi struct {
}

func (m *AdminCarouselApi) CreateCarousel(c *gin.Context) {
	var req manageReq.MallCarouselAddParam
	_ = c.ShouldBindJSON(&req)
	if err := mallAdminCarouselService.CreateCarousel(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMsg("创建失败"+err.Error(), c)
	} else {
		response.OkWithMsg("创建成功", c)
	}
}

func (m *AdminCarouselApi) DeleteCarousel(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallAdminCarouselService.DeleteCarousel(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMsg("删除失败"+err.Error(), c)
	} else {
		response.OkWithMsg("删除成功", c)
	}
}

func (m *AdminCarouselApi) UpdateCarousel(c *gin.Context) {
	var req manageReq.MallCarouselUpdateParam
	_ = c.ShouldBindJSON(&req)
	if err := mallAdminCarouselService.UpdateCarousel(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMsg("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMsg("更新成功", c)
	}
}

// FindCarousel --- FindMallCarousel 用id查询MallCarousel
func (m *AdminCarouselApi) FindCarousel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err, mallCarousel := mallAdminCarouselService.GetCarousel(id); err != nil {
		global.GVA_LOG.Error("查询失败!"+err.Error(), zap.Error(err))
		response.FailWithMsg("查询失败", c)
	} else {
		response.OkWithData(mallCarousel, c)
	}
}

// GetCarouselList --- GetMallCarouselList 分页获取MallCarousel列表
func (m *AdminCarouselApi) GetCarouselList(c *gin.Context) {
	var pageInfo manageReq.MallCarouselSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallAdminCarouselService.GetCarouselInfoList(pageInfo); err != nil {
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
