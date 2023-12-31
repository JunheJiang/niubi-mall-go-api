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

type CarouselApi struct {
}

func (m *CarouselApi) CreateCarousel(c *gin.Context) {
	var req manageReq.MallCarouselAddParam
	_ = c.ShouldBindJSON(&req)
	if err := mallAdminCarouselService.CreateCarousel(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		resp_vo.FailWithMsg("创建失败"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("创建成功", c)
	}
}

func (m *CarouselApi) DeleteCarousel(c *gin.Context) {
	var ids req_param.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallAdminCarouselService.DeleteCarousel(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		resp_vo.FailWithMsg("删除失败"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("删除成功", c)
	}
}

func (m *CarouselApi) UpdateCarousel(c *gin.Context) {
	var req manageReq.MallCarouselUpdateParam
	_ = c.ShouldBindJSON(&req)
	if err := mallAdminCarouselService.UpdateCarousel(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败:"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}
}

// FindCarousel --- FindMallCarousel 用id查询MallCarousel
func (m *CarouselApi) FindCarousel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err, mallCarousel := mallAdminCarouselService.GetCarousel(id); err != nil {
		global.GVA_LOG.Error("查询失败!"+err.Error(), zap.Error(err))
		resp_vo.FailWithMsg("查询失败", c)
	} else {
		resp_vo.OkWithData(mallCarousel, c)
	}
}

// GetCarouselList --- GetMallCarouselList 分页获取MallCarousel列表
func (m *CarouselApi) GetCarouselList(c *gin.Context) {
	var pageInfo manageReq.MallCarouselSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallAdminCarouselService.GetCarouselInfoList(pageInfo); err != nil {
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
