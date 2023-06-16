package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	"niubi-mall/model/common/enum"
	"niubi-mall/model/common/request"
	"niubi-mall/model/common/response"
	manageReq "niubi-mall/model/manage/request"
	manageResponse "niubi-mall/model/manage/response"
	"strconv"
)

type AdminGoodsCategoryApi struct {
}

// CreateCategory 新建商品分类
func (g *AdminGoodsCategoryApi) CreateCategory(c *gin.Context) {
	var category manageReq.MallGoodsCategoryReq
	_ = c.ShouldBindJSON(&category)
	if err := mallAdminGoodsCategoryService.AddCategory(category); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMsg("创建失败:"+err.Error(), c)
	} else {
		response.OkWithMsg("创建成功", c)
	}
}

// UpdateCategory 修改商品分类信息
func (g *AdminGoodsCategoryApi) UpdateCategory(c *gin.Context) {
	var category manageReq.MallGoodsCategoryReq
	_ = c.ShouldBindJSON(&category)
	if err := mallAdminGoodsCategoryService.UpdateCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMsg("更新失败，存在相同分类", c)
	} else {
		response.OkWithMsg("更新成功", c)
	}
}

// GetCategoryList 获取商品分类
func (g *AdminGoodsCategoryApi) GetCategoryList(c *gin.Context) {
	var req manageReq.SearchCategoryParams
	_ = c.ShouldBindQuery(&req)
	if err, list, total := mallAdminGoodsCategoryService.SelectCategoryPage(req); err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMsg("获取失败:"+err.Error(), c)
	} else {
		response.OkWithData(response.PageResult{
			List:       list,
			TotalCount: total,
			CurPage:    req.PageNumber,
			PageSize:   req.PageSize,
			TotalPage:  int(total) / req.PageSize,
		}, c)
	}
}

// GetCategory 通过id获取分类数据
func (g *AdminGoodsCategoryApi) GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsCategory := mallAdminGoodsCategoryService.SelectCategoryById(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMsg("获取失败:"+err.Error(), c)
	} else {
		response.OkWithData(manageResponse.GoodsCategoryResponse{GoodsCategory: goodsCategory}, c)
	}
}

// DelCategory 设置分类失效
func (g *AdminGoodsCategoryApi) DelCategory(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err, _ := mallAdminGoodsCategoryService.DeleteGoodsCategoriesByIds(ids); err != nil {
		global.GVA_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMsg("删除失败"+err.Error(), c)
	} else {
		response.OkWithMsg("删除成功", c)
	}

}

// ListForSelect 用于三级分类联动效果制作
func (g *AdminGoodsCategoryApi) ListForSelect(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, goodsCategory := mallAdminGoodsCategoryService.SelectCategoryById(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMsg("获取失败"+err.Error(), c)
	}
	level := goodsCategory.CategoryLevel
	if level == enum.LevelThree.Code() ||
		level == enum.Default.Code() {
		response.FailWithMsg("参数异常", c)
	}
	categoryResult := make(map[string]interface{})
	if level == enum.LevelOne.Code() {
		_, levelTwoList := mallAdminGoodsCategoryService.SelectByLevelAndParentIdsAndNumber(id, enum.LevelTwo.Code())
		if levelTwoList != nil {
			_, levelThreeList := mallAdminGoodsCategoryService.SelectByLevelAndParentIdsAndNumber(int(levelTwoList[0].CategoryId), enum.LevelThree.Code())
			categoryResult["secondLevelCategories"] = levelTwoList
			categoryResult["thirdLevelCategories"] = levelThreeList
		}
	}
	if level == enum.LevelTwo.Code() {
		_, levelThreeList := mallAdminGoodsCategoryService.SelectByLevelAndParentIdsAndNumber(id, enum.LevelThree.Code())
		categoryResult["thirdLevelCategories"] = levelThreeList
	}
	response.OkWithData(categoryResult, c)
}
