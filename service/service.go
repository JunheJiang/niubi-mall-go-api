package service

import (
	"niubi-mall/service/admin"
	"niubi-mall/service/client"
	"niubi-mall/service/example"
)

type AppServiceGroup struct {
	//has a
	AdminServiceGroup   admin.AdminServiceGroup
	ClientServiceGroup  client.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(AppServiceGroup)
