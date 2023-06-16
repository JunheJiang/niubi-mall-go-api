package request

import "niubi-mall/model/common/req_param"

type MallShopCartSearch struct {
	req_param.PageInfo
}

type SaveCartItemParam struct {
	GoodsCount int `json:"goodsCount"`
	GoodsId    int `json:"goodsId"`
}

type UpdateCartItemParam struct {
	CartItemId int `json:"cartItemId"`
	GoodsCount int `json:"goodsCount"`
}
