package mall

import (
	"niubi-mall/global"
	"niubi-mall/model/mall"
)

type UserTokenService struct {
}

func (m *UserTokenService) ExistUserToken(token string) (err error, mallUserToken mall.UserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallUserToken).Error
	return
}

func (m *UserTokenService) DeleteMallUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]mall.UserToken{}, "token =?", token).Error
	return err
}
