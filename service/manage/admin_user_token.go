package manage

import (
	"niubi-mall/global"
	"niubi-mall/model/manage/db_entity"
)

type AdminUserTokenService struct {
}

func (m *AdminUserTokenService) ExistAdminToken(token string) (err error, mallAdminUserToken db_entity.MallAdminUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&mallAdminUserToken).Error
	return
}

func (m *AdminUserTokenService) DeleteMallAdminUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]db_entity.MallAdminUserToken{}, "token =?", token).Error
	return err
}
