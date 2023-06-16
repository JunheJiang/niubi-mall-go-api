package service

import (
	"niubi-mall/service/example"
	"niubi-mall/service/mall"
	"niubi-mall/service/manage"
)

type AppServiceGroup struct {
	//has a
	AdminServiceGroup   manage.AdminServiceGroup
	MallServiceGroup    mall.ClientServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(AppServiceGroup)
