package response

import "niubi-mall/model/manage"

type GoodsCategoryResponse struct {
	GoodsCategory manage.MallGoodsCategory `json:"mallGoodsCategory"`
}
