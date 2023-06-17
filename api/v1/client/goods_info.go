package client

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	"niubi-mall/model/common/resp_vo"
	"strconv"
)

type GoodsInfoApi struct {
}

// 商品搜索

func (m *GoodsInfoApi) GoodsSearch(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	goodsCategoryId, _ := strconv.Atoi(c.Query("goodsCategoryId"))
	keyword := c.Query("keyword")
	orderBy := c.Query("orderBy")
	if err, list, total := mallGoodsInfoService.MallGoodsListBySearch(pageNumber, goodsCategoryId, keyword, orderBy); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		resp_vo.FailWithMsg("查询失败"+err.Error(), c)
	} else {
		resp_vo.OkWithDetail(resp_vo.PageResult{
			List:       list,
			TotalCount: total,
			CurPage:    pageNumber,
			PageSize:   10,
		}, "获取成功", c)
	}
}

func (m *GoodsInfoApi) GoodsDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsInfo := mallGoodsInfoService.GetMallGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		resp_vo.FailWithMsg("查询失败"+err.Error(), c)
	}
	resp_vo.OkWithData(goodsInfo, c)
}
