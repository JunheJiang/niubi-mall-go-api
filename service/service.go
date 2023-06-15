package service

import (
	"niubi-mall/service/mall"
	"niubi-mall/service/manage"
)

type ServiceGroup struct {
	//has a
	ManageServiceGroup manage.ServiceGroup
	MallServiceGroup   mall.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
