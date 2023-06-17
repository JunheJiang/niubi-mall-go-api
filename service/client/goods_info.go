package client

import (
	"errors"
	"github.com/jinzhu/copier"
	"niubi-mall/global"
	"niubi-mall/model/admin/db_entity"
	response "niubi-mall/model/client/resp_vo"
	"niubi-mall/utils/str"
)

type GoodsInfoService struct {
}

//todo with es
// MallGoodsListBySearch 商品搜索分页

func (m *GoodsInfoService) MallGoodsListBySearch(pageNumber int, goodsCategoryId int,
	keyword string, orderBy string) (err error, searchGoodsList []response.GoodsSearchResponse, total int64) {
	// 根据搜索条件查询
	var goodsList []db_entity.MallGoodsInfo
	db := global.GVA_DB.Model(&db_entity.MallGoodsInfo{})

	if keyword != "" {
		db.Where("goods_name like ? or goods_intro like ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	if goodsCategoryId >= 0 {
		db.Where("goods_category_id= ?", goodsCategoryId)
	}

	err = db.Count(&total).Error

	switch orderBy {
	case "new":
		db.Order("goods_id desc")
	case "price":
		db.Order("selling_price asc")
	default:
		db.Order("stock_num desc")
	}

	limit := 10
	offset := 10 * (pageNumber - 1)
	err = db.Limit(limit).Offset(offset).Find(&goodsList).Error

	// 返回查询结果
	for _, goods := range goodsList {
		searchGoods := response.GoodsSearchResponse{
			GoodsId:       goods.GoodsId,
			GoodsName:     str.SubStrLen(goods.GoodsName, 28),
			GoodsIntro:    str.SubStrLen(goods.GoodsIntro, 28),
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
		}
		searchGoodsList = append(searchGoodsList, searchGoods)
	}
	return
}

// GetMallGoodsInfo 获取商品信息
func (m *GoodsInfoService) GetMallGoodsInfo(id int) (err error, res response.GoodsInfoDetailResponse) {
	var mallGoodsInfo db_entity.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id = ?", id).First(&mallGoodsInfo).Error
	if mallGoodsInfo.GoodsSellStatus != 0 {
		return errors.New("商品已下架"), response.GoodsInfoDetailResponse{}
	}

	err = copier.Copy(&res, &mallGoodsInfo)
	if err != nil {
		return err, response.GoodsInfoDetailResponse{}
	}

	var list []string
	list = append(list, mallGoodsInfo.GoodsCarousel)
	res.GoodsCarouselList = list
	return
}
