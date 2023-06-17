package client

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	request "niubi-mall/model/client/req_param"
	"niubi-mall/model/common/resp_vo"
	"niubi-mall/utils/no"
	"strconv"
)

type ShopCartApi struct {
}

func (m *ShopCartApi) CartItemList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, shopCartItem := mallShopCartService.GetMyShoppingCartItems(token); err != nil {
		global.GVA_LOG.Error("获取购物车失败", zap.Error(err))
		resp_vo.FailWithMsg("获取购物车失败:"+err.Error(), c)
	} else {
		resp_vo.OkWithData(shopCartItem, c)
	}
}

func (m *ShopCartApi) SaveMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	var req request.SaveCartItemParam
	_ = c.ShouldBindJSON(&req)

	if err := mallShopCartService.SaveMallCartItem(token, req); err != nil {
		global.GVA_LOG.Error("添加购物车失败", zap.Error(err))
		resp_vo.FailWithMsg("添加购物车失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("添加购物车成功", c)
}

func (m *ShopCartApi) UpdateMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	var req request.UpdateCartItemParam
	_ = c.ShouldBindJSON(&req)

	if err := mallShopCartService.UpdateMallCartItem(token, req); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		resp_vo.FailWithMsg("修改购物车失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("修改购物车成功", c)
}

func (m *ShopCartApi) DelMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	id, _ := strconv.Atoi(c.Param("newBeeMallShoppingCartItemId"))
	if err := mallShopCartService.DeleteMallCartItem(token, id); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		resp_vo.FailWithMsg("修改购物车失败:"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("修改购物车成功", c)
	}
}

func (m *ShopCartApi) ToSettle(c *gin.Context) {
	cartItemIdsStr := c.Query("cartItemIds")
	token := c.GetHeader("token")

	cartItemIds := no.StrToInt(cartItemIdsStr)
	if err, cartItemRes := mallShopCartService.GetCartItemsForSettle(token, cartItemIds); err != nil {
		global.GVA_LOG.Error("获取购物明细异常：", zap.Error(err))
		resp_vo.FailWithMsg("获取购物明细异常:"+err.Error(), c)
	} else {
		resp_vo.OkWithData(cartItemRes, c)
	}

}
