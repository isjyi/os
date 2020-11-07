package models

import (
	"errors"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/tools/app/msg"
)

type SysDictData struct {
	Id         uint64 `gorm:"primary_key;auto_increment;" json:"id"`
	DictTypeId uint64 `gorm:"" json:"dict_type_id"`                         //字典类型
	Label      string `gorm:"size:128;comment:标签" json:"label"`             //数据标签
	Sort       uint8  `gorm:"comment:排序" json:"sort"`                       //显示顺序
	Value      uint8  `gorm:"comment:数据键值;" json:"value"`                   //数据键值
	Status     uint8  `gorm:"comment:状态 0 正常 1 停用;default:0" json:"status"` //状态
	IsDefault  uint8  `gorm:"comment:默认 ;" json:"is_default"`               // 默认
	CreateBy   uint64 `gorm:"comment:创建者;" json:"create_by"`                //创建者
	UpdateBy   uint64 `gorm:"comment:更新者;" json:"update_by"`                //更新者
	Remark     string `gorm:"size:255;" json:"remark"`                      //备注
	BaseModel

	Params    string `gorm:"-" json:"params"`
	DataScope string `gorm:"-" json:"dataScope"`
} //@name DictData

func (SysDictData) TableName() string {
	return "sys_dict_data"
}

func (e *SysDictData) Get() ([]SysDictData, error) {
	var doc []SysDictData

	table := global.Eloquent.Table(e.TableName())
	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}
	if e.Label != "" {
		table = table.Where("label = ?", e.Label)
	}
	if e.DictTypeId != 0 {
		table = table.Where("dict_type_id = ?", e.DictTypeId)
	}

	result := table.Order("sort").Find(&doc)

	if result.Error != nil {
		return doc, result.Error
	}

	if result.RowsAffected == 0 {
		return doc, errors.New(msg.NotFound)
	}

	return doc, nil
}
