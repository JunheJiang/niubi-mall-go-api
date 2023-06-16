package client

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	"niubi-mall/model/common/resp_vo"
)

type GoodsCategoryApi struct {
}

// GetGoodsCategory ---返回分类数据(首页调用)
func (m *GoodsCategoryApi) GetGoodsCategory(c *gin.Context) {
	err, list := mallGoodsCategoryService.GetCategoriesForIndex()

	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		resp_vo.FailWithMsg("查询失败"+err.Error(), c)
	}
	resp_vo.OkWithData(list, c)
}
