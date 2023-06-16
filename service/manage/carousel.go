package manage

import (
	"errors"
	"gorm.io/gorm"
	"niubi-mall/global"
	"niubi-mall/model/common"
	"niubi-mall/model/common/req_param"
	"niubi-mall/model/manage/db_entity"
	manageReq "niubi-mall/model/manage/req_param"
	"niubi-mall/utils/check"
	"strconv"
	"time"
)

type AdminCarouselService struct {
}

func (m *AdminCarouselService) CreateCarousel(req manageReq.MallCarouselAddParam) (err error) {
	carouseRank, _ := strconv.Atoi(req.CarouselRank)

	mallCarousel := db_entity.MallCarousel{
		CarouselUrl:  req.CarouselUrl,
		RedirectUrl:  req.RedirectUrl,
		CarouselRank: carouseRank,
		CreateTime:   common.JSONTime{Time: time.Now()},
		UpdateTime:   common.JSONTime{Time: time.Now()},
	}

	// 这个校验理论上应该放在api层，但是因为前端的传值是string，而我们的校验规则是Int,所以只能转换格式后再校验
	if err = check.Verify(mallCarousel, check.CarouselAddParamVerify); err != nil {
		return errors.New(err.Error())
	}

	err = global.GVA_DB.Create(&mallCarousel).Error
	return err
}

func (m *AdminCarouselService) DeleteCarousel(ids req_param.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&db_entity.MallCarousel{}, "carousel_id in ?", ids.Ids).Error
	return err
}

func (m *AdminCarouselService) UpdateCarousel(req manageReq.MallCarouselUpdateParam) (err error) {

	if errors.Is(global.GVA_DB.Where("carousel_id = ?", req.CarouselId).First(&db_entity.MallCarousel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("未查询到记录！")
	}

	carouseRank, _ := strconv.Atoi(req.CarouselRank)
	mallCarousel := db_entity.MallCarousel{
		CarouselUrl:  req.CarouselUrl,
		RedirectUrl:  req.RedirectUrl,
		CarouselRank: carouseRank,
		UpdateTime:   common.JSONTime{Time: time.Now()},
	}

	// 这个校验理论上应该放在api层，但是因为前端的传值是string，而我们的校验规则是Int,所以只能转换格式后再校验
	if err = check.Verify(mallCarousel, check.CarouselAddParamVerify); err != nil {
		return errors.New(err.Error())
	}
	err = global.GVA_DB.Where("carousel_id = ?", req.CarouselId).UpdateColumns(&mallCarousel).Error
	return err
}

func (m *AdminCarouselService) GetCarousel(id int) (err error, mallCarousel db_entity.MallCarousel) {
	err = global.GVA_DB.Where("carousel_id = ?", id).First(&mallCarousel).Error
	return
}

func (m *AdminCarouselService) GetCarouselInfoList(info manageReq.MallCarouselSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&db_entity.MallCarousel{})
	var mallCarousels []db_entity.MallCarousel
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("carousel_rank desc").Find(&mallCarousels).Error
	return err, mallCarousels, total
}
