package initialize

import (
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/model"
)

// AutoMigrate run auto migration
func AutoMigrate() {
	db := global.OS_DB

	db.AutoMigrate(model.User{})
	global.OS_LOG.Debug("register table success")
}
