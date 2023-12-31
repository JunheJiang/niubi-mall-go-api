package admin

import (
	"errors"
	"gorm.io/gorm"
	"niubi-mall/global"
	"niubi-mall/model/admin/db_entity"
	manageReq "niubi-mall/model/admin/req_param"
	"niubi-mall/model/common"
	"niubi-mall/model/common/req_param"
	"niubi-mall/utils/check"
	"niubi-mall/utils/no"
	"strconv"
	"time"
)

type AdminGoodsCategoryService struct {
}

// AddCategory 添加商品分类
func (m *AdminGoodsCategoryService) AddCategory(req manageReq.MallGoodsCategoryReq) (err error) {

	if !errors.Is(global.GVA_DB.Where("category_level=? AND category_name=? AND is_deleted=0",
		req.CategoryLevel, req.CategoryName).First(&db_entity.MallGoodsCategory{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同分类")
	}

	rank, _ := strconv.Atoi(req.CategoryRank)
	category := db_entity.MallGoodsCategory{
		CategoryLevel: req.CategoryLevel,
		CategoryName:  req.CategoryName,
		CategoryRank:  rank,
		IsDeleted:     0,
		CreateTime:    common.JSONTime{Time: time.Now()},
		UpdateTime:    common.JSONTime{Time: time.Now()},
	}

	// 这个校验理论上应该放在api层，但是因为前端的传值是string，而我们的校验规则是Int,所以只能转换格式后再校验
	if err = check.Verify(category, check.GoodsCategoryVerify); err != nil {
		return errors.New(err.Error())
	}
	return global.GVA_DB.Create(&category).Error
}

// UpdateCategory 更新商品分类
func (m *AdminGoodsCategoryService) UpdateCategory(req manageReq.MallGoodsCategoryReq) (err error) {

	if !errors.Is(global.GVA_DB.Where("category_level=? AND category_name=? AND is_deleted=0",
		req.CategoryLevel, req.CategoryName).First(&db_entity.MallGoodsCategory{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同分类")
	}

	rank, _ := strconv.Atoi(req.CategoryRank)
	category := db_entity.MallGoodsCategory{
		CategoryName: req.CategoryName,
		CategoryRank: rank,
		UpdateTime:   common.JSONTime{Time: time.Now()},
	}

	// 这个校验理论上应该放在api层，但是因为前端的传值是string，而我们的校验规则是Int,所以只能转换格式后再校验
	if err := check.Verify(category, check.GoodsCategoryVerify); err != nil {
		return errors.New(err.Error())
	}
	return global.GVA_DB.Where("category_id =?", req.CategoryId).Updates(&category).Error

}

// SelectCategoryPage 获取分类分页数据
func (m *AdminGoodsCategoryService) SelectCategoryPage(req manageReq.SearchCategoryParams) (err error, list interface{}, total int64) {
	limit := req.PageSize
	if limit > 1000 {
		limit = 1000
	}
	offset := req.PageSize * (req.PageNumber - 1)
	db := global.GVA_DB.Model(&db_entity.MallGoodsCategory{})
	var categoryList []db_entity.MallGoodsCategory

	if no.NumInList(req.CategoryLevel, []int{1, 2, 3}) {
		db.Where("category_level=?", req.CategoryLevel)
	}

	if req.ParentId >= 0 {
		db.Where("parent_id=?", req.ParentId)
	}

	err = db.Where("is_deleted=0").Count(&total).Error

	if err != nil {
		return err, categoryList, total

	} else {
		db = db.Where("is_deleted=0").Order("category_rank desc").Limit(limit).Offset(offset)
		err = db.Find(&categoryList).Error
	}
	return err, categoryList, total
}

// SelectCategoryById 获取单个分类数据
func (m *AdminGoodsCategoryService) SelectCategoryById(categoryId int) (err error, goodsCategory db_entity.MallGoodsCategory) {
	err = global.GVA_DB.Where("category_id=?", categoryId).First(&goodsCategory).Error
	return err, goodsCategory
}

// DeleteGoodsCategoriesByIds 批量设置失效
func (m *AdminGoodsCategoryService) DeleteGoodsCategoriesByIds(ids req_param.IdsReq) (err error, goodsCategory db_entity.MallGoodsCategory) {
	err = global.GVA_DB.Where("category_id in ?", ids.Ids).UpdateColumns(db_entity.MallGoodsCategory{IsDeleted: 1}).Error
	return err, goodsCategory
}

func (m *AdminGoodsCategoryService) SelectByLevelAndParentIdsAndNumber(parentId int, categoryLevel int) (err error, goodsCategories []db_entity.MallGoodsCategory) {
	err = global.GVA_DB.Where("category_id in ?", parentId).Where("category_level=?", categoryLevel).Where("is_deleted=0").Error
	return err, goodsCategories
}
