package client

import (
	"errors"
	"github.com/jinzhu/copier"
	"niubi-mall/global"
	amdin "niubi-mall/model/admin/db_entity"
	manageReq "niubi-mall/model/admin/req_param"
	client "niubi-mall/model/client/db_entity"
	response "niubi-mall/model/client/resp_vo"
	"niubi-mall/model/common"
	"niubi-mall/model/common/enum"
	"niubi-mall/utils/no"
	"time"
)

type OrderService struct {
}

// SaveOrder 保存订单
func (m *OrderService) SaveOrder(token string, userAddress client.MallUserAddress,
	myShoppingCartItems []response.CartItemResponse) (err error, orderNo string) {
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), orderNo
	}
	var itemIdList []int
	var goodsIds []int
	for _, cartItem := range myShoppingCartItems {
		itemIdList = append(itemIdList, cartItem.CartItemId)
		goodsIds = append(goodsIds, cartItem.GoodsId)
	}
	var newBeeMallGoods []amdin.MallGoodsInfo

	global.GVA_DB.Where("goods_id in ? ", goodsIds).Find(&newBeeMallGoods)
	//检查是否包含已下架商品
	for _, mallGoods := range newBeeMallGoods {
		if mallGoods.GoodsSellStatus != enum.GOODS_UNDER.Code() {
			return errors.New("已下架，无法生成订单"), orderNo
		}
	}

	newBeeMallGoodsMap := make(map[int]amdin.MallGoodsInfo)
	for _, mallGoods := range newBeeMallGoods {
		newBeeMallGoodsMap[mallGoods.GoodsId] = mallGoods
	}
	//判断商品库存
	for _, shoppingCartItemVO := range myShoppingCartItems {
		//查出的商品中不存在购物车中的这条关联商品数据，直接返回错误提醒
		if _, ok := newBeeMallGoodsMap[shoppingCartItemVO.GoodsId]; !ok {
			return errors.New("购物车数据异常！"), orderNo
		}
		if shoppingCartItemVO.GoodsCount > newBeeMallGoodsMap[shoppingCartItemVO.GoodsId].StockNum {
			return errors.New("库存不足！"), orderNo
		}
	}
	//删除购物项
	if len(itemIdList) > 0 && len(goodsIds) > 0 {
		if err = global.GVA_DB.Where("cart_item_id in ?", itemIdList).Updates(client.MallShoppingCartItem{IsDeleted: 1}).Error; err == nil {
			var stockNumDTOS []manageReq.StockNumDTO
			copier.Copy(&stockNumDTOS, &myShoppingCartItems)
			for _, stockNumDTO := range stockNumDTOS {
				var goodsInfo amdin.MallGoodsInfo
				global.GVA_DB.Where("goods_id =?", stockNumDTO.GoodsId).First(&goodsInfo)
				if err = global.GVA_DB.Where("goods_id =? and stock_num>= ? and goods_sell_status = 0", stockNumDTO.GoodsId, stockNumDTO.GoodsCount).Updates(amdin.MallGoodsInfo{StockNum: goodsInfo.StockNum - stockNumDTO.GoodsCount}).Error; err != nil {
					return errors.New("库存不足！"), orderNo
				}
			}
			//生成订单号
			orderNo = no.GenOrderNo()
			priceTotal := 0
			//保存订单
			var newBeeMallOrder amdin.MallOrder
			newBeeMallOrder.OrderNo = orderNo
			newBeeMallOrder.UserId = userToken.UserId
			//总价
			for _, newBeeMallShoppingCartItemVO := range myShoppingCartItems {
				priceTotal = priceTotal + newBeeMallShoppingCartItemVO.GoodsCount*newBeeMallShoppingCartItemVO.SellingPrice
			}
			if priceTotal < 1 {
				return errors.New("订单价格异常！"), orderNo
			}
			newBeeMallOrder.CreateTime = common.JSONTime{Time: time.Now()}
			newBeeMallOrder.UpdateTime = common.JSONTime{Time: time.Now()}
			newBeeMallOrder.TotalPrice = priceTotal
			newBeeMallOrder.ExtraInfo = ""
			//生成订单项并保存订单项纪录
			if err = global.GVA_DB.Save(&newBeeMallOrder).Error; err != nil {
				return errors.New("订单入库失败！"), orderNo
			}
			//生成订单收货地址快照，并保存至数据库
			var newBeeMallOrderAddress client.MallOrderAddress
			copier.Copy(&newBeeMallOrderAddress, &userAddress)
			newBeeMallOrderAddress.OrderId = newBeeMallOrder.OrderId
			//生成所有的订单项快照，并保存至数据库
			var newBeeMallOrderItems []amdin.MallOrderItem
			for _, newBeeMallShoppingCartItemVO := range myShoppingCartItems {
				var newBeeMallOrderItem amdin.MallOrderItem
				copier.Copy(&newBeeMallOrderItem, &newBeeMallShoppingCartItemVO)
				newBeeMallOrderItem.OrderId = newBeeMallOrder.OrderId
				newBeeMallOrderItem.CreateTime = common.JSONTime{Time: time.Now()}
				newBeeMallOrderItems = append(newBeeMallOrderItems, newBeeMallOrderItem)
			}
			if err = global.GVA_DB.Save(&newBeeMallOrderItems).Error; err != nil {
				return err, orderNo
			}
		}
	}
	return
}

// PaySuccess 支付订单
func (m *OrderService) PaySuccess(orderNo string, payType int) (err error) {
	var mallOrder amdin.MallOrder
	err = global.GVA_DB.Where("order_no = ? and is_deleted=0 ", orderNo).First(&mallOrder).Error
	if mallOrder != (amdin.MallOrder{}) {
		if mallOrder.OrderStatus != 0 {
			return errors.New("订单状态异常！")
		}
		mallOrder.OrderStatus = enum.ORDER_PAID.Code()
		mallOrder.PayType = payType
		mallOrder.PayStatus = 1
		mallOrder.PayTime = common.JSONTime{time.Now()}
		mallOrder.UpdateTime = common.JSONTime{time.Now()}
		err = global.GVA_DB.Save(&mallOrder).Error
	}
	return
}

// FinishOrder 完结订单
func (m *OrderService) FinishOrder(token string, orderNo string) (err error) {
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var mallOrder amdin.MallOrder
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return errors.New("未查询到记录！")
	}
	if mallOrder.UserId != userToken.UserId {
		return errors.New("禁止该操作！")
	}
	mallOrder.OrderStatus = enum.ORDER_SUCCESS.Code()
	mallOrder.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&mallOrder).Error
	return
}

// CancelOrder 关闭订单
func (m *OrderService) CancelOrder(token string, orderNo string) (err error) {
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var mallOrder amdin.MallOrder
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return errors.New("未查询到记录！")
	}

	if mallOrder.UserId != userToken.UserId {
		return errors.New("禁止该操作！")
	}

	if no.NumInList(mallOrder.OrderStatus,
		[]int{enum.ORDER_SUCCESS.Code(),
			enum.ORDER_CLOSED_BY_MALLUSER.Code(),
			enum.ORDER_CLOSED_BY_EXPIRED.Code(),
			enum.ORDER_CLOSED_BY_JUDGE.Code()}) {
		return errors.New("订单状态异常！")
	}

	mallOrder.OrderStatus = enum.ORDER_CLOSED_BY_MALLUSER.Code()
	mallOrder.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&mallOrder).Error
	return
}

// GetOrderDetailByOrderNo 获取订单详情
func (m *OrderService) GetOrderDetailByOrderNo(token string, orderNo string) (err error, orderDetail response.MallOrderDetailVO) {
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), orderDetail
	}
	var mallOrder amdin.MallOrder
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return errors.New("未查询到记录！"), orderDetail
	}
	if mallOrder.UserId != userToken.UserId {
		return errors.New("禁止该操作！"), orderDetail
	}
	var orderItems []amdin.MallOrderItem
	err = global.GVA_DB.Where("order_id = ?", mallOrder.OrderId).Find(&orderItems).Error
	if len(orderItems) <= 0 {
		return errors.New("订单项不存在！"), orderDetail
	}

	var newBeeMallOrderItemVOS []response.NewBeeMallOrderItemVO
	copier.Copy(&newBeeMallOrderItemVOS, &orderItems)
	copier.Copy(&orderDetail, &mallOrder)
	// 订单状态前端显示为中文
	_, OrderStatusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(orderDetail.OrderStatus)
	_, payTapStr := enum.GetNewBeeMallOrderStatusEnumByStatus(orderDetail.PayType)
	orderDetail.OrderStatusString = OrderStatusStr
	orderDetail.PayTypeString = payTapStr
	orderDetail.NewBeeMallOrderItemVOS = newBeeMallOrderItemVOS

	return
}

// MallOrderListBySearch 搜索订单
func (m *OrderService) MallOrderListBySearch(token string, pageNumber int, status string) (err error, list []response.MallOrderResponse, total int64) {
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), list, total
	}
	// 根据搜索条件查询
	var newBeeMallOrders []amdin.MallOrder
	db := global.GVA_DB.Model(&newBeeMallOrders)

	if status != "" {
		db.Where("order_status = ?", status)
	}
	err = db.Where("user_id =? and is_deleted=0 ", userToken.UserId).Count(&total).Error
	//这里前段没有做滚动加载，直接显示全部订单
	//limit := 5
	offset := 5 * (pageNumber - 1)
	err = db.Offset(offset).Order(" order_id desc").Find(&newBeeMallOrders).Error

	var orderListVOS []response.MallOrderResponse
	if total > 0 {
		//数据转换 将实体类转成vo
		copier.Copy(&orderListVOS, &newBeeMallOrders)
		//设置订单状态中文显示值
		for _, newBeeMallOrderListVO := range orderListVOS {
			_, statusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderListVO.OrderStatus)
			newBeeMallOrderListVO.OrderStatusString = statusStr
		}
		// 返回订单id
		var orderIds []int
		for _, order := range newBeeMallOrders {
			orderIds = append(orderIds, order.OrderId)
		}
		//获取OrderItem
		var orderItems []amdin.MallOrderItem
		if len(orderIds) > 0 {
			global.GVA_DB.Where("order_id in ?", orderIds).Find(&orderItems)
			itemByOrderIdMap := make(map[int][]amdin.MallOrderItem)
			for _, orderItem := range orderItems {
				itemByOrderIdMap[orderItem.OrderId] = []amdin.MallOrderItem{}
			}
			for k, v := range itemByOrderIdMap {
				for _, orderItem := range orderItems {
					if k == orderItem.OrderId {
						v = append(v, orderItem)
					}
					itemByOrderIdMap[k] = v
				}
			}
			//封装每个订单列表对象的订单项数据
			for _, newBeeMallOrderListVO := range orderListVOS {
				if _, ok := itemByOrderIdMap[newBeeMallOrderListVO.OrderId]; ok {
					orderItemListTemp := itemByOrderIdMap[newBeeMallOrderListVO.OrderId]
					var newBeeMallOrderItemVOS []response.NewBeeMallOrderItemVO
					copier.Copy(&newBeeMallOrderItemVOS, &orderItemListTemp)
					newBeeMallOrderListVO.NewBeeMallOrderItemVOS = newBeeMallOrderItemVOS
					_, OrderStatusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderListVO.OrderStatus)
					newBeeMallOrderListVO.OrderStatusString = OrderStatusStr
					list = append(list, newBeeMallOrderListVO)
				}
			}
		}
	}
	return err, list, total
}
