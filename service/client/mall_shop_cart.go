package client

import (
	"errors"
	"github.com/jinzhu/copier"
	"niubi-mall/global"
	amdin "niubi-mall/model/admin/db_entity"
	client "niubi-mall/model/client/db_entity"
	request "niubi-mall/model/client/req_param"
	response "niubi-mall/model/client/resp_vo"
	"niubi-mall/model/common"
	"niubi-mall/utils/str"
	"time"
)

type ShopCartService struct {
}

// GetMyShoppingCartItems 不分页
func (m *ShopCartService) GetMyShoppingCartItems(token string) (err error, cartItems []response.CartItemResponse) {
	var userToken client.MallUserToken
	var shopCartItems []client.MallShoppingCartItem
	var goodsInfos []amdin.MallGoodsInfo

	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), cartItems
	}

	global.GVA_DB.Where("user_id=? and is_deleted = 0", userToken.UserId).Find(&shopCartItems)
	var goodsIds []int
	for _, shopCartItem := range shopCartItems {
		goodsIds = append(goodsIds, shopCartItem.GoodsId)
	}
	global.GVA_DB.Where("goods_id in ?", goodsIds).Find(&goodsInfos)
	goodsMap := make(map[int]amdin.MallGoodsInfo)
	for _, goodsInfo := range goodsInfos {
		goodsMap[goodsInfo.GoodsId] = goodsInfo
	}

	for _, v := range shopCartItems {
		var cartItem response.CartItemResponse
		copier.Copy(&cartItem, &v)

		if _, ok := goodsMap[v.GoodsId]; ok {
			goodsInfo := goodsMap[v.GoodsId]
			cartItem.GoodsName = goodsInfo.GoodsName
			cartItem.GoodsCoverImg = goodsInfo.GoodsCoverImg
			cartItem.SellingPrice = goodsInfo.SellingPrice
		}
		cartItems = append(cartItems, cartItem)
	}
	return
}

func (m *ShopCartService) SaveMallCartItem(token string, req request.SaveCartItemParam) (err error) {
	if req.GoodsCount < 1 {
		return errors.New("商品数量不能小于 1 ！")

	}
	if req.GoodsCount > 5 {
		return errors.New("超出单个商品的最大购买数量！")
	}
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var shopCartItems []client.MallShoppingCartItem
	// 是否已存在商品
	err = global.GVA_DB.Where("user_id = ? and goods_id = ? and is_deleted = 0", userToken.UserId, req.GoodsId).Find(&shopCartItems).Error
	if err != nil {
		return errors.New("已存在！无需重复添加！")
	}

	err = global.GVA_DB.Where("goods_id = ? ", req.GoodsId).First(&amdin.MallGoodsInfo{}).Error
	if err != nil {
		return errors.New(" 商品为空")
	}

	var total int64
	global.GVA_DB.Where("user_id =?  and is_deleted = 0", userToken.UserId).Count(&total)
	if total > 20 {
		return errors.New("超出购物车最大容量！")
	}

	var shopCartItem client.MallShoppingCartItem
	if err = copier.Copy(&shopCartItem, &req); err != nil {
		return err
	}
	shopCartItem.UserId = userToken.UserId
	shopCartItem.CreateTime = common.JSONTime{Time: time.Now()}
	shopCartItem.UpdateTime = common.JSONTime{Time: time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *ShopCartService) UpdateMallCartItem(token string, req request.UpdateCartItemParam) (err error) {
	//超出单个商品的最大数量
	if req.GoodsCount > 5 {
		return errors.New("超出单个商品的最大购买数量！")
	}
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var shopCartItem client.MallShoppingCartItem
	if err = global.GVA_DB.Where("cart_item_id=? and is_deleted = 0", req.CartItemId).First(&shopCartItem).Error; err != nil {
		return errors.New("未查询到记录！")
	}

	if shopCartItem.UserId != userToken.UserId {
		return errors.New("禁止该操作！")
	}
	shopCartItem.GoodsCount = req.GoodsCount
	shopCartItem.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *ShopCartService) DeleteMallCartItem(token string, id int) (err error) {
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var shopCartItem client.MallShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).First(&shopCartItem).Error
	if err != nil {
		return
	}

	if userToken.UserId != shopCartItem.UserId {
		return errors.New("禁止该操作！")
	}
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).UpdateColumns(&client.MallShoppingCartItem{IsDeleted: 1}).Error
	return
}

func (m *ShopCartService) GetCartItemsForSettle(token string, cartItemIds []int) (err error, cartItemRes []response.CartItemResponse) {
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), cartItemRes
	}

	var shopCartItems []client.MallShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id in (?) and user_id = ? and is_deleted = 0", cartItemIds, userToken.UserId).Find(&shopCartItems).Error
	if err != nil {
		return
	}
	_, cartItemRes = getMallShoppingCartItemVOS(shopCartItems)
	//购物车算价
	priceTotal := 0
	for _, cartItem := range cartItemRes {
		priceTotal = priceTotal + cartItem.GoodsCount*cartItem.SellingPrice
	}
	return
}

// 购物车数据转换
func getMallShoppingCartItemVOS(cartItems []client.MallShoppingCartItem) (err error, cartItemsRes []response.CartItemResponse) {
	var goodsIds []int
	for _, cartItem := range cartItems {
		goodsIds = append(goodsIds, cartItem.GoodsId)
	}

	var newBeeMallGoods []amdin.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id in ?", goodsIds).Find(&newBeeMallGoods).Error
	if err != nil {
		return
	}

	newBeeMallGoodsMap := make(map[int]amdin.MallGoodsInfo)
	for _, goodsInfo := range newBeeMallGoods {
		newBeeMallGoodsMap[goodsInfo.GoodsId] = goodsInfo
	}

	for _, cartItem := range cartItems {
		var cartItemRes response.CartItemResponse
		copier.Copy(&cartItemRes, &cartItem)
		// 是否包含key
		if _, ok := newBeeMallGoodsMap[cartItemRes.GoodsId]; ok {
			newBeeMallGoodsTemp := newBeeMallGoodsMap[cartItemRes.GoodsId]
			cartItemRes.GoodsCoverImg = newBeeMallGoodsTemp.GoodsCoverImg
			goodsName := str.SubStrLen(newBeeMallGoodsTemp.GoodsName, 28)
			cartItemRes.GoodsName = goodsName
			cartItemRes.SellingPrice = newBeeMallGoodsTemp.SellingPrice
			cartItemsRes = append(cartItemsRes, cartItemRes)
		}
	}
	return
}
