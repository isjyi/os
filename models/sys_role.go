package models

type SysRole struct {
	ID       uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	RoleName string `json:"role_name" gorm:"size:50;comment:'角色名称'"`               // 角色名称
	Status   bool   `json:"status" gorm:"size:1;comment:'状态 1 正常 0 删除';default:1"` // 状态
	// RoleKey  string `json:"roleKey" gorm:"size:128;comment:'角色唯一标识'"`              // 角色唯一标识
	CreateBy uint64 `json:"create_by" gorm:"comment:'创建者'"` // 创建者
	UpdateBy uint64 `json:"update_by" gorm:"comment:'修改者'"` // 修改者
	Remark   string `json:"remark" gorm:"comment:'备注'"`     //备注
	Params   string `json:"params" gorm:"-"`
	MenuIds  []int  `json:"menu_ids" gorm:"-"`
	DeptIds  []int  `json:"dept_ids" gorm:"-"`
	BaseModel
}

func (SysRole) TableName() string {
	return "sys_role"
}
