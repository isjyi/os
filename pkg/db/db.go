package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func New() (db *gorm.DB, err error) {

	db, err = gorm.Open("mysql", dsn())

	if err != nil {
		return
	}

	if os.Getenv("DEBUG") != "" {
		db.LogMode(true)
	}
	db.DB().SetConnMaxLifetime(time.Second)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)
	DB = db
	return
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
		viper.GetString("db.charset"),
		viper.GetBool("db.parseTime"),
		viper.GetString("db.loc"),
	)
}
