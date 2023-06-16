package mall

import (
	"niubi-mall/global"
	"niubi-mall/model/mall/db_entity"
)

type UserTokenService struct {
}

func (m *UserTokenService) ExistUserToken(token string) (err error, mallUserToken db_entity.UserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallUserToken).Error
	return
}

func (m *UserTokenService) DeleteMallUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]db_entity.UserToken{}, "token =?", token).Error
	return err
}
