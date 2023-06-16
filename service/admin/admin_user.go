package admin

import (
	"errors"
	"gorm.io/gorm"
	"niubi-mall/global"
	"niubi-mall/model/admin/db_entity"
	"niubi-mall/model/admin/req_param"
	"niubi-mall/utils/data"
	"niubi-mall/utils/no"
	"strconv"
	"strings"
	"time"
)

type AdminUserService struct {
}

func (m *AdminUserService) CreateMallAdminUser(user db_entity.MallAdminUser) (err error) {
	err = global.GVA_DB.Where("login_user_name = ?", user.LoginUserName).First(&db_entity.MallAdminUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("存在相同用户名")
	}

	err = global.GVA_DB.Create(&user).Error
	return err
}

func (m *AdminUserService) UpdateMallAdminName(token string, param req_param.MallUpdateNameParam) (err error) {
	var adminUserToken db_entity.MallAdminUserToken
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	err = global.GVA_DB.Where("admin_user_id = ?", adminUserToken.AdminUserId).Updates(&db_entity.MallAdminUser{
		LoginUserName: param.LoginUserName,
		NickName:      param.NickName,
	}).Error
	return err
}

func (m *AdminUserService) UpdateMallAdminPassWord(token string, req req_param.MallUpdatePasswordParam) (err error) {
	var adminUserToken db_entity.MallAdminUserToken
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("用户未登录")
	}
	var adminUser db_entity.MallAdminUser
	err = global.GVA_DB.Where("admin_user_id =?", adminUserToken.AdminUserId).First(&adminUser).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	if adminUser.LoginPassword != req.OriginalPassword {
		return errors.New("原密码不正确")
	}
	adminUser.LoginPassword = req.NewPassword

	err = global.GVA_DB.Where("admin_user_id=?", adminUser.AdminUserId).Updates(&adminUser).Error
	return
}

// GetMallAdminUser 根据id获取MallAdminUser记录
func (m *AdminUserService) GetMallAdminUser(token string) (err error, mallAdminUser db_entity.MallAdminUser) {
	var adminToken db_entity.MallAdminUserToken
	if errors.Is(global.GVA_DB.Where("token =?", token).First(&adminToken).Error, gorm.ErrRecordNotFound) {
		return errors.New("不存在的用户"), mallAdminUser
	}
	err = global.GVA_DB.Where("admin_user_id = ?", adminToken.AdminUserId).First(&mallAdminUser).Error
	return err, mallAdminUser
}

// AdminLogin 管理员登陆
func (m *AdminUserService) AdminLogin(params req_param.MallAdminLoginParam) (err error, mallAdminUser db_entity.MallAdminUser, adminToken db_entity.MallAdminUserToken) {
	err = global.GVA_DB.Where("login_user_name=? AND login_password=?", params.UserName, params.PasswordMd5).First(&mallAdminUser).Error
	if mallAdminUser != (db_entity.MallAdminUser{}) {
		token := getNewToken(time.Now().UnixNano()/1e6, int(mallAdminUser.AdminUserId))
		global.GVA_DB.Where("admin_user_id", mallAdminUser.AdminUserId).First(&adminToken)
		nowDate := time.Now()
		// 48小时过期
		expireTime, _ := time.ParseDuration("48h")
		expireDate := nowDate.Add(expireTime)
		// 没有token新增，有token 则更新
		if adminToken == (db_entity.MallAdminUserToken{}) {
			adminToken.AdminUserId = mallAdminUser.AdminUserId
			adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if err = global.GVA_DB.Create(&adminToken).Error; err != nil {
				return
			}
		} else {
			adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&adminToken).Error; err != nil {
				return
			}
		}
	}
	return err, mallAdminUser, adminToken

}

func getNewToken(timeInt int64, userId int) (token string) {
	var build strings.Builder
	build.WriteString(strconv.FormatInt(timeInt, 10))
	build.WriteString(strconv.Itoa(userId))
	build.WriteString(no.GenValidateCode(6))
	return data.MD5V([]byte(build.String()))
}
