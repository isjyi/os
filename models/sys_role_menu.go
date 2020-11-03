package models

type SysRoleMenu struct {
	RoleId uint64 `gorm:"comment:角色id;index" json:"role_id"`
	MenuId uint64 `gorm:"comment:菜单id;index" json:"menu_id"`
	BaseModel
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
