package models

import "github.com/isjyi/os/global"

type SysSetting struct {
	ID   uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"type:varchar(256);"`
	Logo string `json:"logo" gorm:"type:varchar(256);"`
	BaseModel
}

func (SysSetting) TableName() string {
	return "sys_setting"
}

func (s *SysSetting) Get() (r SysSetting, e error) {
	res := global.Eloquent.Take(&r)
	e = res.Error
	return
}
