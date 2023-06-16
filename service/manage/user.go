package manage

import (
	"errors"
	"niubi-mall/global"
	"niubi-mall/model/common/req_param"
	"niubi-mall/model/manage/db_entity"
	manageReq "niubi-mall/model/manage/req_param"
)

type UserService struct {
}

func (m *UserService) LockUser(id req_param.IdsReq, lockStatus int) (err error) {
	if lockStatus != 0 && lockStatus != 1 {
		return errors.New("操作非法")
	}
	err = global.GVA_DB.Model(&db_entity.MallUser{}).Where("user_id in ?", id.Ids).Update("lock_flag", lockStatus).Error
	return err
}

func (m *UserService) GetMallUserInfoList(info manageReq.MallUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	db := global.GVA_DB.Model(&db_entity.MallUser{})

	var mallUsers []db_entity.MallUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&mallUsers).Error
	return err, mallUsers, total
}
