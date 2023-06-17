package client

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"niubi-mall/global"
	request "niubi-mall/model/client/req_param"
	"niubi-mall/model/common/resp_vo"
	"strconv"
)

type UserAddressApi struct {
}

func (m *UserAddressApi) AddressList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddressList := mallUserAddressService.GetMyAddress(token); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		resp_vo.FailWithMsg("获取地址失败:"+err.Error(), c)
	} else if len(userAddressList) == 0 {
		resp_vo.OkWithData(nil, c)
	} else {
		resp_vo.OkWithData(userAddressList, c)
	}
}

func (m *UserAddressApi) SaveUserAddress(c *gin.Context) {
	var req request.AddAddressParam
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	err := mallUserAddressService.SaveUserAddress(token, req)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		resp_vo.FailWithMsg("创建失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("创建成功", c)
}

func (m *UserAddressApi) UpdateMallUserAddress(c *gin.Context) {
	var req request.UpdateAddressParam
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	err := mallUserAddressService.UpdateUserAddress(token, req)
	if err != nil {
		global.GVA_LOG.Error("更新用户地址失败", zap.Error(err))
		resp_vo.FailWithMsg("更新用户地址失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("更新用户地址成功", c)
}

func (m *UserAddressApi) GetMallUserAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("addressId"))
	token := c.GetHeader("token")
	if err, userAddress := mallUserAddressService.GetMallUserAddressById(token, id); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		resp_vo.FailWithMsg("获取地址失败:"+err.Error(), c)
	} else {
		resp_vo.OkWithData(userAddress, c)
	}
}

func (m *UserAddressApi) GetMallUserDefaultAddress(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddress := mallUserAddressService.GetMallUserDefaultAddress(token); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		resp_vo.FailWithMsg("获取地址失败:"+err.Error(), c)
	} else {
		resp_vo.OkWithData(userAddress, c)
	}
}

func (m *UserAddressApi) DeleteUserAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("addressId"))
	token := c.GetHeader("token")
	err := mallUserAddressService.DeleteUserAddress(token, id)
	if err != nil {
		global.GVA_LOG.Error("删除用户地址失败", zap.Error(err))
		resp_vo.FailWithMsg("删除用户地址失败:"+err.Error(), c)
	}
	resp_vo.OkWithMsg("删除用户地址成功", c)
}
