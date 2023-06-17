package client

import (
	"errors"
	"github.com/jinzhu/copier"
	"niubi-mall/global"
	client "niubi-mall/model/client/db_entity"
	request "niubi-mall/model/client/req_param"
	"niubi-mall/model/common"
	"time"
)

type UserAddressService struct {
}

// GetMyAddress 获取收货地址
func (m *UserAddressService) GetMyAddress(token string) (err error, userAddress []client.MallUserAddress) {
	var userToken client.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), userAddress
	}

	global.GVA_DB.Where("user_id=? and is_deleted=0", userToken.UserId).Find(&userAddress)
	return
}

// SaveUserAddress 保存用户地址
func (m *UserAddressService) SaveUserAddress(token string, req request.AddAddressParam) (err error) {
	var userToken client.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}
	// 是否新增了默认地址，将之前的默认地址设置为非默认
	var defaultAddress client.MallUserAddress
	copier.Copy(&defaultAddress, &req)
	defaultAddress.CreateTime = common.JSONTime{Time: time.Now()}
	defaultAddress.UpdateTime = common.JSONTime{Time: time.Now()}
	defaultAddress.UserId = userToken.UserId

	if req.DefaultFlag == 1 {
		global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userToken.UserId).First(&defaultAddress)
		if defaultAddress != (client.MallUserAddress{}) {
			defaultAddress.UpdateTime = common.JSONTime{time.Now()}
			err = global.GVA_DB.Save(&defaultAddress).Error
			if err != nil {
				return
			}
		}
	} else {
		err = global.GVA_DB.Create(&defaultAddress).Error
		if err != nil {
			return
		}
	}
	return
}

// UpdateUserAddress 更新用户地址
func (m *UserAddressService) UpdateUserAddress(token string, req request.UpdateAddressParam) (err error) {
	var userToken client.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}
	var userAddress client.MallUserAddress
	if err = global.GVA_DB.Where("address_id =? and user_id =?", req.AddressId, userToken.UserId).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！")
	}
	if req.DefaultFlag == 1 {
		var defaultUserAddress client.MallUserAddress
		global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userToken.UserId).First(&defaultUserAddress)
		if defaultUserAddress != (client.MallUserAddress{}) {
			defaultUserAddress.DefaultFlag = 0
			defaultUserAddress.UpdateTime = common.JSONTime{time.Now()}
			err = global.GVA_DB.Save(&defaultUserAddress).Error
			if err != nil {
				return
			}
		}
	}
	err = copier.Copy(&userAddress, &req)
	if err != nil {
		return
	}
	userAddress.UpdateTime = common.JSONTime{time.Now()}
	userAddress.UserId = userToken.UserId
	err = global.GVA_DB.Save(&userAddress).Error
	return
}

func (m *UserAddressService) GetMallUserAddressById(token string, id int) (err error, userAddress client.MallUserAddress) {
	var userToken client.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址"), userAddress
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！"), userAddress
	}
	return
}

func (m *UserAddressService) GetMallUserDefaultAddress(token string) (err error, userAddress client.MallUserAddress) {
	var userToken client.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	if err = global.GVA_DB.Where("user_id =? and default_flag =1 and is_deleted = 0 ", userToken.UserId).First(&userAddress).Error; err != nil {
		return errors.New("不存在默认地址失败"), userAddress
	}
	return
}

func (m *UserAddressService) DeleteUserAddress(token string, id int) (err error) {
	var userToken client.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}

	var userAddress client.MallUserAddress
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！")
	}
	err = global.GVA_DB.Delete(&userAddress).Error
	return

}
