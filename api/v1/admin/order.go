package admin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	"niubi-mall/model/common/req_param"
	"niubi-mall/model/common/resp_vo"
)

type OrderApi struct {
}

// CheckDoneOrder 发货
func (m *OrderApi) CheckDoneOrder(c *gin.Context) {
	var IDS req_param.IdsReq
	_ = c.ShouldBindJSON(&IDS)

	if err := mallAdminOrderService.CheckDone(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败", c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}
}

// CheckOutOrder 出库
func (m *OrderApi) CheckOutOrder(c *gin.Context) {
	var IDS req_param.IdsReq
	_ = c.ShouldBindJSON(&IDS)

	if err := mallAdminOrderService.CheckOut(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败", c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}
}

// CloseOrder 关闭订单
func (m *OrderApi) CloseOrder(c *gin.Context) {
	var IDS req_param.IdsReq
	_ = c.ShouldBindJSON(&IDS)

	if err := mallAdminOrderService.CloseOrder(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败", c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}
}

// FindMallOrder 用id查询MallOrder
func (m *OrderApi) FindMallOrder(c *gin.Context) {
	id := c.Param("orderId")

	if err, newBeeMallOrderDetailVO := mallAdminOrderService.GetMallOrder(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		resp_vo.FailWithMsg("查询失败", c)
	} else {
		resp_vo.OkWithData(newBeeMallOrderDetailVO, c)
	}
}

// GetMallOrderList 分页获取MallOrder列表
func (m *OrderApi) GetMallOrderList(c *gin.Context) {
	var pageInfo req_param.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)

	orderNo := c.Query("orderNo")
	orderStatus := c.Query("orderStatus")

	if err, list, total := mallAdminOrderService.GetMallOrderInfoList(pageInfo, orderNo, orderStatus); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
