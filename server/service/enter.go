package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/cat_management"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	Cat_managementServiceGroup cat_management.ServiceGroup
}
