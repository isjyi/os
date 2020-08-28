package initialize

import (
	"os"
	"time"

	"github.com/isjyi/os/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

func Mysql() {
	c := global.OS_CONFIG.Mysql
	db, err := gorm.Open("mysql", c.DNS())

	if err != nil {
		global.OS_LOG.Error("MYSQL启动异常", zap.Error(err))
		os.Exit(0)
	}

	global.OS_DB = db

	if global.OS_CONFIG.System.Mode == "debug" {
		global.OS_DB.LogMode(c.LogMode)
	}

	global.OS_DB.DB().SetConnMaxLifetime(time.Second * time.Duration(global.OS_CONFIG.Mysql.MaxLifetime))
	global.OS_DB.DB().SetMaxIdleConns(global.OS_CONFIG.Mysql.MaxIdleConns)
	global.OS_DB.DB().SetMaxOpenConns(global.OS_CONFIG.Mysql.MaxOpenConns)
}
