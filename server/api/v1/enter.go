package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/cat_management"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup         system.ApiGroup
	ExampleApiGroup        example.ApiGroup
	Cat_managementApiGroup cat_management.ApiGroup
}
