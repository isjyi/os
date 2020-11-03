package models

type BaseModel struct {
	CreatedAt int32 `json:"createdAt"`
	UpdatedAt int32 `json:"updatedAt"`
	DeletedAt int32 `json:"deletedAt" gorm:"default:null" sql:"index"`
}
