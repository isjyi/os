package models

type SysDictType struct {
	Id       uint64 `gorm:"primary_key;auto_increment;" json:"id"`
	Name     string `gorm:"size:128;comment:字典名称" json:"name"` //字典名称
	Type     string `gorm:"size:128;" json:"type"`             //字典类型
	Status   uint8  `gorm:"comment:状态 0正常 1停用;" json:"status"` //状态
	CreateBy uint64 `gorm:"comment:创建者;" json:"create_by"`     //创建者
	UpdateBy uint64 `gorm:"comment:更新者;" json:"update_by"`     //更新者
	Remark   string `gorm:"size:255;comment:备注" json:"remark"` //备注

	Datas []SysDictData `gorm:"foreignKey:DictTypeId" json:"datas"` //
	BaseModel
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
