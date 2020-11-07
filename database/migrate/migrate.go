package migrate

import (
	"github.com/isjyi/os/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.SysUser{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.SysRole{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Casbin{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.SysSetting{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.SysRole{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.SysRoleMenu{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Menu{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.SysDictType{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.SysDictData{})
	if err != nil {
		return err
	}

	// models.DataInit()
	return err
}

func CustomMigrate(db *gorm.DB, table interface{}) error {
	return db.AutoMigrate(&table)
}
