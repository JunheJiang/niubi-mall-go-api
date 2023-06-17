package client

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	request "niubi-mall/model/client/req_param"
	"niubi-mall/model/common/resp_vo"
	"niubi-mall/utils/check"
	"strconv"
)

type OrderApi struct {
}

func (m *OrderApi) SaveOrder(c *gin.Context) {
	var saveOrderParam request.SaveOrderParam
	_ = c.ShouldBindJSON(&saveOrderParam)
	if err := check.Verify(saveOrderParam, check.SaveOrderParamVerify); err != nil {
		resp_vo.FailWithMsg(err.Error(), c)
	}
	token := c.GetHeader("token")

	priceTotal := 0
	err, itemsForSave := mallShopCartService.GetCartItemsForSettle(token, saveOrderParam.CartItemIds)
	if len(itemsForSave) < 1 {
		resp_vo.FailWithMsg("无数据:"+err.Error(), c)
	} else {
		//总价
		for _, newBeeMallShoppingCartItemVO := range itemsForSave {
			priceTotal = priceTotal + newBeeMallShoppingCartItemVO.GoodsCount*newBeeMallShoppingCartItemVO.SellingPrice
		}
		if priceTotal < 1 {
			resp_vo.FailWithMsg("价格异常", c)
		}
		_, userAddress := mallUserAddressService.GetMallUserDefaultAddress(token)
		if err, saveOrderResult := mallOrderService.SaveOrder(token, userAddress, itemsForSave); err != nil {
			global.GVA_LOG.Error("生成订单失败", zap.Error(err))
			resp_vo.FailWithMsg("生成订单失败:"+err.Error(), c)
		} else {
			resp_vo.OkWithData(saveOrderResult, c)
		}
	}
}

func (m *OrderApi) PaySuccess(c *gin.Context) {
	orderNo := c.Query("orderNo")
	payType, _ := strconv.Atoi(c.Query("payType"))
	if err := mallOrderService.PaySuccess(orderNo, payType); err != nil {
		global.GVA_LOG.Error("订单支付失败", zap.Error(err))
		resp_vo.FailWithMsg("订单支付失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("订单支付成功", c)
}

func (m *OrderApi) FinishOrder(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	if err := mallOrderService.FinishOrder(token, orderNo); err != nil {
		global.GVA_LOG.Error("订单签收失败", zap.Error(err))
		resp_vo.FailWithMsg("订单签收失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("订单签收成功", c)
}

func (m *OrderApi) CancelOrder(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	if err := mallOrderService.CancelOrder(token, orderNo); err != nil {
		global.GVA_LOG.Error("订单签收失败", zap.Error(err))
		resp_vo.FailWithMsg("订单签收失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("订单签收成功", c)
}

func (m *OrderApi) OrderDetailPage(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	if err, orderDetail := mallOrderService.GetOrderDetailByOrderNo(token, orderNo); err != nil {
		global.GVA_LOG.Error("查询订单详情接口", zap.Error(err))
		resp_vo.FailWithMsg("查询订单详情接口:"+err.Error(), c)
	} else {
		resp_vo.OkWithData(orderDetail, c)
	}
}

func (m *OrderApi) OrderList(c *gin.Context) {
	token := c.GetHeader("token")
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	status := c.Query("status")
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if err, list, total := mallOrderService.MallOrderListBySearch(token, pageNumber, status); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		resp_vo.FailWithMsg("查询失败"+err.Error(), c)
	} else if len(list) < 1 {
		// 前端项目这里有一个取数逻辑，如果数组为空，数组需要为[] 不能是Null
		resp_vo.OkWithDetail(resp_vo.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurPage:    pageNumber,
			PageSize:   5,
		}, "SUCCESS", c)
	} else {
		resp_vo.OkWithDetail(resp_vo.PageResult{
			List:       list,
			TotalCount: total,
			CurPage:    pageNumber,
			PageSize:   5,
		}, "SUCCESS", c)
	}

}
