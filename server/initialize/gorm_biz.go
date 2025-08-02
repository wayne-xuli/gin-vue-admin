package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cat_management"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(cat_management.CatInfo{})
	if err != nil {
		return err
	}
	return nil
}
