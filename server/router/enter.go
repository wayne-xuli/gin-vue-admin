package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/cat_management"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System         system.RouterGroup
	Example        example.RouterGroup
	Cat_management cat_management.RouterGroup
}
