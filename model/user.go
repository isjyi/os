package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Phone    int    `gorm:"unique;not null;index:idx_users_phone" json:"phone"`
	Password string `gorm:"size 64" json:"password"`
}
