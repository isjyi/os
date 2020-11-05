package models

import "github.com/isjyi/os/global"

type SysRoleMenu struct {
	RoleId uint64 `gorm:"comment:角色id;index" json:"role_id"`
	MenuId uint64 `gorm:"comment:菜单id;index" json:"menu_id"`
	BaseModel
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

func (rm *SysRoleMenu) GetPermits() ([]string, error) {
	var r []Menu
	table := global.Eloquent.Select("sys_menu.permission").Table("sys_menu").Joins("left join sys_role_menu on sys_menu.id = sys_role_menu.menu_id")

	table = table.Where("role_id = ?", rm.RoleId)

	table = table.Where("sys_menu.menu_type in(?,?)", 2, 3)

	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	var list []string
	for i := 0; i < len(r); i++ {
		if r[i].Permission != "" {
			list = append(list, r[i].Permission)
		}
	}
	return list, nil
}
