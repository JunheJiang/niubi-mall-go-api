package service

import (
	"niubi-mall/service/example"
	"niubi-mall/service/mall"
	"niubi-mall/service/manage"
)

type ServiceGroup struct {
	//has a
	AdminServiceGroup   manage.ServiceGroup
	MallServiceGroup    mall.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
