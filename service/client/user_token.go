package client

import (
	"niubi-mall/global"
	"niubi-mall/model/client/db_entity"
)

type UserTokenService struct {
}

func (m *UserTokenService) ExistUserToken(token string) (err error, mallUserToken db_entity.MallUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallUserToken).Error
	return
}

func (m *UserTokenService) DeleteMallUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]db_entity.MallUserToken{}, "token =?", token).Error
	return err
}
