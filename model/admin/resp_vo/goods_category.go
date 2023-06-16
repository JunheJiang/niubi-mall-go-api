package resp_vo

import (
	"niubi-mall/model/admin/db_entity"
)

type GoodsCategoryResponse struct {
	GoodsCategory db_entity.MallGoodsCategory `json:"mallGoodsCategory"`
}
