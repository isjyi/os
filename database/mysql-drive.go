package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Mysql struct {
}

func (m *Mysql) Setup() {
	var err error
	db, err := sql.Open("mysql", m.GetConnect())

	if err != nil {
		global.Logger.Fatal(tools.Red(m.GetDriver()+" connect error :"), zap.Error(err))
	}

	global.Eloquent, err = m.Open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
	})
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

	if config.OSConfig.Logger.EnabledDB {
		global.Eloquent.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		})
	}
}

// 打开数据库连接
func (m *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// 获取数据库连接
func (m *Mysql) GetConnect() string {
	return config.OSConfig.Database.Source
}

func (m *Mysql) GetDriver() string {
	return config.OSConfig.Database.Driver
}
