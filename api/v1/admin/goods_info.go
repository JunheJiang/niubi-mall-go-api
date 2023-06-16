package admin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	"niubi-mall/model/admin/db_entity"
	manageReq "niubi-mall/model/admin/req_param"
	"niubi-mall/model/common/req_param"
	"niubi-mall/model/common/resp_vo"
	"strconv"
)

type AdminGoodsInfoApi struct {
}

func (m *AdminGoodsInfoApi) CreateGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manageReq.GoodsInfoAddParam
	_ = c.ShouldBindJSON(&mallGoodsInfo)

	if err := mallAdminGoodsInfoService.CreateMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		resp_vo.FailWithMsg("创建失败!"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("创建成功", c)
	}
}

// DeleteGoodsInfo --- DeleteMallGoodsInfo 删除MallGoodsInfo
func (m *AdminGoodsInfoApi) DeleteGoodsInfo(c *gin.Context) {
	var mallGoodsInfo db_entity.MallGoodsInfo
	_ = c.ShouldBindJSON(&mallGoodsInfo)

	if err := mallAdminGoodsInfoService.DeleteMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		resp_vo.FailWithMsg("删除失败"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("删除成功", c)
	}
}

// ChangeMallGoodsInfoByIds 批量删除MallGoodsInfo

func (m *AdminGoodsInfoApi) ChangeGoodsInfoByIds(c *gin.Context) {
	var IDS req_param.IdsReq
	_ = c.ShouldBindJSON(&IDS)

	sellStatus := c.Param("status")

	if err := mallAdminGoodsInfoService.ChangeMallGoodsInfoByIds(IDS, sellStatus); err != nil {
		global.GVA_LOG.Error("修改商品状态失败!", zap.Error(err))
		resp_vo.FailWithMsg("修改商品状态失败"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("修改商品状态成功", c)
	}
}

// UpdateMallGoodsInfo 更新MallGoodsInfo

func (m *AdminGoodsInfoApi) UpdateGoodsInfo(c *gin.Context) {
	var mallGoodsInfo manageReq.GoodsInfoUpdateParam
	_ = c.ShouldBindJSON(&mallGoodsInfo)

	if err := mallAdminGoodsInfoService.UpdateMallGoodsInfo(mallGoodsInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		resp_vo.FailWithMsg("更新失败"+err.Error(), c)
	} else {
		resp_vo.OkWithMsg("更新成功", c)
	}
}

// FindMallGoodsInfo 用id查询MallGoodsInfo

func (m *AdminGoodsInfoApi) FindGoodsInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsInfo := mallAdminGoodsInfoService.GetMallGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		resp_vo.FailWithMsg("查询失败"+err.Error(), c)
	}
	goodsInfoRes := make(map[string]interface{})

	goodsInfoRes["goods"] = goodsInfo

	if _, thirdCategory := mallAdminGoodsCategoryService.SelectCategoryById(goodsInfo.GoodsCategoryId); thirdCategory != (db_entity.MallGoodsCategory{}) {
		goodsInfoRes["thirdCategory"] = thirdCategory
		if _, secondCategory := mallAdminGoodsCategoryService.SelectCategoryById(thirdCategory.ParentId); secondCategory != (db_entity.MallGoodsCategory{}) {
			goodsInfoRes["secondCategory"] = secondCategory
			if _, firstCategory := mallAdminGoodsCategoryService.SelectCategoryById(secondCategory.ParentId); firstCategory != (db_entity.MallGoodsCategory{}) {
				goodsInfoRes["firstCategory"] = firstCategory
			}
		}
	}
	resp_vo.OkWithData(goodsInfoRes, c)

}

// GetMallGoodsInfoList 分页获取MallGoodsInfo列表

func (m *AdminGoodsInfoApi) GetGoodsInfoList(c *gin.Context) {
	var pageInfo manageReq.MallGoodsInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)

	goodsName := c.Query("goodsName")
	goodsSellStatus := c.Query("goodsSellStatus")

	if err, list, total := mallAdminGoodsInfoService.GetMallGoodsInfoInfoList(pageInfo, goodsName, goodsSellStatus); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		resp_vo.FailWithMsg("获取失败"+err.Error(), c)
	} else {
		resp_vo.OkWithDetail(resp_vo.PageResult{
			List:       list,
			TotalCount: total,
			CurPage:    pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
