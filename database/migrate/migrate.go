package migrate

import (
	"github.com/isjyi/os/models"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	db.SingularTable(true)
	err := db.AutoMigrate(&models.SysUser{}).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.SysRole{}).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.CasbinRule{}).Error
	if err != nil {
		return err
	}

	// models.DataInit()
	return err
}

func CustomMigrate(db *gorm.DB, table interface{}) error {
	db.SingularTable(true)
	return db.AutoMigrate(&table).Error
}
