package client

import (
	"niubi-mall/global"
	"niubi-mall/model/admin/db_entity"
	response "niubi-mall/model/client/resp_vo"
	"niubi-mall/utils/str"
)

type IndexInfoService struct {
}

// GetConfigGoodsForIndex 首页返回相关IndexConfig
func (m *IndexInfoService) GetConfigGoodsForIndex(configType int, num int) (err error, list interface{}) {
	var indexConfigs []db_entity.MallIndexConfig
	err = global.GVA_DB.Where("config_type = ?", configType).
		Where("is_deleted = 0").
		Order("config_rank desc").
		Limit(num).Find(&indexConfigs).Error
	if err != nil {
		return
	}

	// 获取商品id
	var ids []int
	for _, indexConfig := range indexConfigs {
		ids = append(ids, indexConfig.GoodsId)
	}

	// 获取商品信息
	var goodsList []db_entity.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id in ?", ids).Find(&goodsList).Error

	var indexGoodsList []response.MallIndexConfigGoodsResponse
	// 超出30个字符显示....
	for _, indexGoods := range goodsList {
		res := response.MallIndexConfigGoodsResponse{
			GoodsId:       indexGoods.GoodsId,
			GoodsName:     str.SubStrLen(indexGoods.GoodsName, 30),
			GoodsIntro:    str.SubStrLen(indexGoods.GoodsIntro, 30),
			GoodsCoverImg: indexGoods.GoodsCoverImg,
			SellingPrice:  indexGoods.SellingPrice,
			Tag:           indexGoods.Tag,
		}
		indexGoodsList = append(indexGoodsList, res)
	}
	return err, indexGoodsList
}
