package client

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"niubi-mall/global"
	"niubi-mall/model/admin/db_entity"
	mall_db_entity "niubi-mall/model/client/db_entity"
	request "niubi-mall/model/client/req_param"
	response "niubi-mall/model/client/resp_vo"
	"niubi-mall/model/common"
	"niubi-mall/utils/data"
	"niubi-mall/utils/no"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
}

// RegisterUser 注册用户
func (m *UserService) RegisterUser(req request.RegisterUserParam) (err error) {
	if !errors.Is(global.GVA_DB.Where("login_name =?", req.LoginName).First(&mall_db_entity.MallUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同用户名")
	}

	return global.GVA_DB.Create(&mall_db_entity.MallUser{
		LoginName:     req.LoginName,
		PasswordMd5:   data.MD5V([]byte(req.Password)),
		IntroduceSign: "随新所欲，蜂富多彩",
		CreateTime:    common.JSONTime{Time: time.Now()},
	}).Error

}

func (m *UserService) UpdateUserInfo(token string, req request.UpdateUserInfoParam) (err error) {
	var userToken mall_db_entity.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}

	var userInfo db_entity.MallUser
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error

	// 若密码为空字符，则表明用户不打算修改密码，使用原密码保存
	if !(req.PasswordMd5 == "") {
		userInfo.PasswordMd5 = data.MD5V([]byte(req.PasswordMd5))
	}
	userInfo.NickName = req.NickName
	userInfo.IntroduceSign = req.IntroduceSign
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).UpdateColumns(&userInfo).Error
	return
}

func (m *UserService) GetUserDetail(token string) (err error, userDetail response.MallUserDetailResponse) {
	var userToken mall_db_entity.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), userDetail
	}

	var userInfo db_entity.MallUser
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error
	if err != nil {
		return errors.New("用户信息获取失败"), userDetail
	}
	err = copier.Copy(&userDetail, &userInfo)
	return
}

func (m *UserService) UserLogin(params request.UserLoginParam) (err error, user db_entity.MallUser, userToken mall_db_entity.MallUserToken) {

	err = global.GVA_DB.Where("login_name=? AND password_md5=?", params.LoginName, params.PasswordMd5).First(&user).Error

	if user != (db_entity.MallUser{}) {
		token := getNewToken(time.Now().UnixNano()/1e6, int(user.UserId))
		global.GVA_DB.Where("user_id", user.UserId).First(&token)
		nowDate := time.Now()
		// 48小时过期
		expireTime, _ := time.ParseDuration("48h")
		expireDate := nowDate.Add(expireTime)
		// 没有token新增，有token 则更新
		if userToken == (mall_db_entity.MallUserToken{}) {
			userToken.UserId = user.UserId
			userToken.Token = token
			userToken.UpdateTime = nowDate
			userToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&userToken).Error; err != nil {
				return
			}
		} else {
			userToken.Token = token
			userToken.UpdateTime = nowDate
			userToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&userToken).Error; err != nil {
				return
			}
		}
	}
	return err, user, userToken
}

func getNewToken(timeInt int64, userId int) (token string) {
	var build strings.Builder
	build.WriteString(strconv.FormatInt(timeInt, 10))
	build.WriteString(strconv.Itoa(userId))
	build.WriteString(no.GenValidateCode(6))
	return data.MD5V([]byte(build.String()))
}
