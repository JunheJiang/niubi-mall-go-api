package client

import (
	"github.com/jinzhu/copier"
	"niubi-mall/global"
	"niubi-mall/model/admin/db_entity"
	response "niubi-mall/model/client/resp_vo"
	"niubi-mall/model/common/enum"
)

type GoodsCategoryService struct {
}

// GetCategoriesForIndex --- todo opt
func (m *GoodsCategoryService) GetCategoriesForIndex() (err error, newBeeMallIndexCategoryVOS []response.NewBeeMallIndexCategoryVO) {

	//获取一级分类的固定数量的数据
	_, firstLevelCategories := selectByLevelAndParentIdsAndNumber([]int{0}, enum.LevelOne.Code(), 20)

	if firstLevelCategories != nil {
		var firstLevelCategoryIds []int

		for _, firstLevelCategory := range firstLevelCategories {
			//append len string make
			firstLevelCategoryIds = append(firstLevelCategoryIds, firstLevelCategory.CategoryId)
		}

		//获取二级分类的数据
		_, secondLevelCategories := selectByLevelAndParentIdsAndNumber(firstLevelCategoryIds, enum.LevelTwo.Code(), 20)
		if secondLevelCategories != nil {

			var secondLevelCategoryIds []int

			for _, secondLevelCategory := range secondLevelCategories {
				secondLevelCategoryIds = append(secondLevelCategoryIds, secondLevelCategory.CategoryId)
			}

			//获取三级分类的数据
			_, thirdLevelCategories := selectByLevelAndParentIdsAndNumber(secondLevelCategoryIds, enum.LevelThree.Code(), 20)
			if thirdLevelCategories != nil {

				//根据 parentId 将 thirdLevelCategories 分组
				thirdLevelCategoryMap := make(map[int][]db_entity.MallGoodsCategory)

				for _, thirdLevelCategory := range thirdLevelCategories {
					thirdLevelCategoryMap[thirdLevelCategory.ParentId] = []db_entity.MallGoodsCategory{}
				}

				for k, v := range thirdLevelCategoryMap {
					for _, third := range thirdLevelCategories {
						if k == third.ParentId {
							v = append(v, third)
						}
						thirdLevelCategoryMap[k] = v
					}
				}

				var secondLevelCategoryVOS []response.SecondLevelCategoryVO

				//处理二级分类
				for _, secondLevelCategory := range secondLevelCategories {

					var secondLevelCategoryVO response.SecondLevelCategoryVO
					err = copier.Copy(&secondLevelCategoryVO, &secondLevelCategory)

					//如果该二级分类下有数据则放入 secondLevelCategoryVOS 对象中
					if _, ok := thirdLevelCategoryMap[secondLevelCategory.CategoryId]; ok {
						//根据二级分类的id取出thirdLevelCategoryMap分组中的三级分类list
						tempGoodsCategories := thirdLevelCategoryMap[secondLevelCategory.CategoryId]

						var thirdLevelCategoryRes []response.ThirdLevelCategoryVO
						err = copier.Copy(&thirdLevelCategoryRes, &tempGoodsCategories)

						secondLevelCategoryVO.ThirdLevelCategoryVOS = thirdLevelCategoryRes
						secondLevelCategoryVOS = append(secondLevelCategoryVOS, secondLevelCategoryVO)
					}

				}

				//处理二级分类
				if secondLevelCategoryVOS != nil {
					//根据 parentId 将 thirdLevelCategories 分组
					secondLevelCategoryVOMap := make(map[int][]response.SecondLevelCategoryVO)

					for _, secondLevelCategory := range secondLevelCategoryVOS {
						secondLevelCategoryVOMap[secondLevelCategory.ParentId] = []response.SecondLevelCategoryVO{}
					}

					for k, v := range secondLevelCategoryVOMap {
						for _, second := range secondLevelCategoryVOS {
							if k == second.ParentId {
								var secondLevelCategory response.SecondLevelCategoryVO
								copier.Copy(&secondLevelCategory, &second)
								v = append(v, secondLevelCategory)
							}
							secondLevelCategoryVOMap[k] = v
						}
					}

					for _, firstCategory := range firstLevelCategories {
						var newBeeMallIndexCategoryVO response.NewBeeMallIndexCategoryVO
						err = copier.Copy(&newBeeMallIndexCategoryVO, &firstCategory)

						//如果该一级分类下有数据则放入 newBeeMallIndexCategoryVOS 对象中
						if _, ok := secondLevelCategoryVOMap[firstCategory.CategoryId]; ok {
							//根据一级分类的id取出secondLevelCategoryVOMap分组中的二级级分类list
							tempGoodsCategories := secondLevelCategoryVOMap[firstCategory.CategoryId]
							newBeeMallIndexCategoryVO.SecondLevelCategoryVOS = tempGoodsCategories
							newBeeMallIndexCategoryVOS = append(newBeeMallIndexCategoryVOS, newBeeMallIndexCategoryVO)
						}
					}
				}
			}
		}
	}
	return
}

// 获取分类数据
func selectByLevelAndParentIdsAndNumber(ids []int, level int, limit int) (err error, categories []db_entity.MallGoodsCategory) {

	err = global.GVA_DB.Where("parent_id in ? and category_level =? and is_deleted = 0", ids, level).
		Order("category_rank desc").Limit(limit).Find(&categories).Error
	return
}
