package models

import "gorm.io/gorm"

type BaseModel struct {
	CreatedAt int32          `json:"createdAt"`
	UpdatedAt int32          `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"default:null" sql:"index"`
}
