package database

import (
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

type Mysql struct {
}

func (m *Mysql) Setup() {
	var err error
	global.Eloquent, err = m.Open()
	if err != nil {
		global.Logger.Fatal(tools.Red(m.GetDriver()+" connect error :"), zap.Error(err))
	} else {
		global.Logger.Info(tools.Green(m.GetDriver() + " connect success !"))
	}
	if global.Eloquent.Error != nil {
		global.Logger.Fatal(tools.Red(" database error :"), zap.Error(global.Eloquent.Error))
	}
	// 数据表前缀
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "os_" + defaultTableName
	// }

	global.Eloquent.LogMode(config.OSConfig.Logger.EnabledDB)
}

// 打开数据库连接
func (m *Mysql) Open() (db *gorm.DB, err error) {
	return gorm.Open(m.GetDriver(), m.GetConnect())
}

// 获取数据库连接
func (m *Mysql) GetConnect() string {
	return config.OSConfig.Database.Source
}

func (m *Mysql) GetDriver() string {
	return config.OSConfig.Database.Driver
}
