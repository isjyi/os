package database

import "github.com/jinzhu/gorm"

type Database interface {
	Setup()
	Open() (db *gorm.DB, err error)
	GetConnect() string
	GetDriver() string
}
