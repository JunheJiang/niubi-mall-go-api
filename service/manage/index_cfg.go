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
	"niubi-mall/utils/no"
	"strconv"
	"time"
)

type AdminIndexConfigService struct {
}

// CreateMallIndexConfig 创建MallIndexConfig记录
func (m *AdminIndexConfigService) CreateMallIndexConfig(req manageReq.MallIndexConfigAddParams) (err error) {
	var goodsInfo db_entity.MallGoodsInfo

	if errors.Is(global.GVA_DB.Where("goods_id=?", req.GoodsId).First(&goodsInfo).Error, gorm.ErrRecordNotFound) {
		return errors.New("商品不存在")
	}

	if errors.Is(global.GVA_DB.Where("config_type =? and goods_id=? and is_deleted=0", req.ConfigType, req.GoodsId).First(&db_entity.MallIndexConfig{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在相同的首页配置项")
	}

	goodsId, _ := strconv.Atoi(req.GoodsId)
	configRank, _ := strconv.Atoi(req.ConfigRank)
	mallIndexConfig := db_entity.MallIndexConfig{
		ConfigName:  req.ConfigName,
		ConfigType:  req.ConfigType,
		GoodsId:     goodsId,
		RedirectUrl: req.RedirectUrl,
		ConfigRank:  configRank,
		CreateTime:  common.JSONTime{Time: time.Now()},
		UpdateTime:  common.JSONTime{Time: time.Now()},
	}

	if err = check.Verify(mallIndexConfig, check.IndexConfigAddParamVerify); err != nil {
		return errors.New(err.Error())
	}

	err = global.GVA_DB.Create(&mallIndexConfig).Error
	return err
}

// DeleteMallIndexConfig 删除MallIndexConfig记录
func (m *AdminIndexConfigService) DeleteMallIndexConfig(ids req_param.IdsReq) (err error) {
	err = global.GVA_DB.Where("config_id in ?", ids.Ids).Delete(&db_entity.MallIndexConfig{}).Error
	return err
}

// UpdateMallIndexConfig 更新MallIndexConfig记录
func (m *AdminIndexConfigService) UpdateMallIndexConfig(req manageReq.MallIndexConfigUpdateParams) (err error) {
	//更新indexConfig
	if errors.Is(global.GVA_DB.Where("goods_id = ?", req.GoodsId).First(&db_entity.MallGoodsInfo{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("商品不存在！")
	}

	if errors.Is(global.GVA_DB.Where("config_id=?", req.ConfigId).First(&db_entity.MallIndexConfig{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("未查询到记录！")
	}

	configRank, _ := strconv.Atoi(req.ConfigRank)
	mallIndexConfig := db_entity.MallIndexConfig{
		ConfigId:    req.ConfigId,
		ConfigType:  req.ConfigType,
		ConfigName:  req.ConfigName,
		RedirectUrl: req.RedirectUrl,
		GoodsId:     req.GoodsId,
		ConfigRank:  configRank,
		UpdateTime:  common.JSONTime{Time: time.Now()},
	}

	if err = check.Verify(mallIndexConfig, check.IndexConfigUpdateParamVerify); err != nil {
		return errors.New(err.Error())
	}

	var newIndexConfig db_entity.MallIndexConfig
	err = global.GVA_DB.Where("config_type=? and goods_id=?", mallIndexConfig.ConfigType, mallIndexConfig.GoodsId).First(&newIndexConfig).Error
	if err != nil && newIndexConfig.ConfigId == mallIndexConfig.ConfigId {
		return errors.New("已存在相同的首页配置项")
	}
	err = global.GVA_DB.Where("config_id=?", mallIndexConfig.ConfigId).Updates(&mallIndexConfig).Error
	return err
}

// GetMallIndexConfig 根据id获取MallIndexConfig记录
func (m *AdminIndexConfigService) GetMallIndexConfig(id uint) (err error, mallIndexConfig db_entity.MallIndexConfig) {
	err = global.GVA_DB.Where("config_id = ?", id).First(&mallIndexConfig).Error
	return
}

// GetMallIndexConfigInfoList 分页获取MallIndexConfig记录
func (m *AdminIndexConfigService) GetMallIndexConfigInfoList(info manageReq.MallIndexConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&db_entity.MallIndexConfig{})
	// todo 有没有更好的方式实现？
	if no.NumInList(info.ConfigType, []int{1, 2, 3, 4, 5}) {
		db.Where("config_type=?", info.ConfigType)
	}

	var mallIndexConfigs []db_entity.MallIndexConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("config_rank desc").Find(&mallIndexConfigs).Error
	return err, mallIndexConfigs, total
}
