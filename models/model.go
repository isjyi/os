package models

type BaseModel struct {
	CreatedAt int `json:"createdAt"`
	UpdatedAt int `json:"updatedAt"`
	DeletedAt int `json:"deletedAt" gorm:"default:null" sql:"index"`
}
