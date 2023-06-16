package manage

import (
	"errors"
	"github.com/jinzhu/copier"
	"niubi-mall/global"
	"niubi-mall/model/common"
	"niubi-mall/model/common/enum"
	"niubi-mall/model/common/req_param"
	"niubi-mall/model/manage/db_entity"
	"niubi-mall/model/manage/resp_vo"
	"strconv"
	"time"
)

type AdminOrderService struct {
}

// CheckDone 修改订单状态为配货成功
func (m *AdminOrderService) CheckDone(ids req_param.IdsReq) (err error) {
	var orders []db_entity.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error

	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {

			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}

			if order.OrderStatus != enum.ORDER_PAID.Code() {
				errorOrders = order.OrderNo + " "
			}

		}

		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).
				UpdateColumns(db_entity.MallOrder{OrderStatus: 2, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单的状态不是支付成功无法执行出库操作")
		}
	}
	return
}

// CheckOut 出库
func (m *AdminOrderService) CheckOut(ids req_param.IdsReq) (err error) {
	var orders []db_entity.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error

	var errorOrders string
	if len(orders) != 0 {

		for _, order := range orders {

			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}

			if order.OrderStatus != enum.ORDER_PAID.Code() && order.OrderStatus != enum.ORDER_PACKAGED.Code() {
				errorOrders = order.OrderNo + " "
			}
		}

		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).
				UpdateColumns(db_entity.MallOrder{OrderStatus: 3, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单的状态不是支付成功或配货完成无法执行出库操作")
		}
	}
	return
}

// CloseOrder 商家关闭订单
func (m *AdminOrderService) CloseOrder(ids req_param.IdsReq) (err error) {
	var orders []db_entity.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error

	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {

			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}

			if order.OrderStatus == enum.ORDER_SUCCESS.Code() || order.OrderStatus < 0 {
				errorOrders = order.OrderNo + " "
			}
		}

		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).
				UpdateColumns(db_entity.MallOrder{OrderStatus: enum.ORDER_CLOSED_BY_JUDGE.Code(), UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单不能执行关闭操作")
		}
	}
	return
}

// GetMallOrder 根据id获取MallOrder记录
func (m *AdminOrderService) GetMallOrder(id string) (err error, newBeeMallOrderDetailVO resp_vo.NewBeeMallOrderDetailVO) {
	var newBeeMallOrder db_entity.MallOrder
	if err = global.GVA_DB.Where("order_id = ?", id).First(&newBeeMallOrder).Error; err != nil {
		//newBeeMallOrderDetailVO ---> nil
		return
	}

	var orderItems []db_entity.MallOrderItem
	if err = global.GVA_DB.Where("order_id = ?", newBeeMallOrder.OrderId).Find(&orderItems).Error; err != nil {
		//newBeeMallOrderDetailVO ---> nil
		return
	}

	//获取订单项数据
	if len(orderItems) > 0 {
		var newBeeMallOrderItemVOS []resp_vo.NewBeeMallOrderItemVO
		copier.Copy(&newBeeMallOrderItemVOS, &orderItems)
		copier.Copy(&newBeeMallOrderDetailVO, &newBeeMallOrder)

		_, OrderStatusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderDetailVO.OrderStatus)
		_, payTapStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderDetailVO.PayType)
		newBeeMallOrderDetailVO.OrderStatusString = OrderStatusStr
		newBeeMallOrderDetailVO.PayTypeString = payTapStr
		newBeeMallOrderDetailVO.NewBeeMallOrderItemVOS = newBeeMallOrderItemVOS
	}
	// err vo
	return
}

// GetMallOrderInfoList 分页获取MallOrder记录
func (m *AdminOrderService) GetMallOrderInfoList(info req_param.PageInfo, orderNo string, orderStatus string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&db_entity.MallOrder{})

	if orderNo != "" {
		db.Where("order_no", orderNo)
	}

	// 0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭
	if orderStatus != "" {
		status, _ := strconv.Atoi(orderStatus)
		db.Where("order_status", status)
	}

	var mallOrders []db_entity.MallOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("update_time desc").Find(&mallOrders).Error
	return err, mallOrders, total
}
