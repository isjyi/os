package mysql

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

func New(logger *zap.Logger) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return
	}

	db.DB().SetConnMaxLifetime(time.Second)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)

	return
}
