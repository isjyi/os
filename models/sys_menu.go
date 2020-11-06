package models

import "github.com/isjyi/os/global"

type Menu struct {
	ID           uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`         // id
	Name         string `json:"name" gorm:"size:128;comment:菜单名称;"`           // 菜单名称
	Title        string `json:"title" gorm:"size:128;comment:菜单标题;"`          // 菜单标题
	Icon         string `json:"icon" gorm:"size:128;comment:菜单图标;"`           // 菜单图标
	Path         string `json:"path" gorm:"size:128;comment:路由地址;"`           // 路由地址
	Paths        string `json:"paths" gorm:"size:50;comment:菜单图标;"`           //层级顺序
	MenuType     uint8  `json:"menu_type" gorm:"comment:菜单类型;"`               //菜单类型
	Action       string `json:"action" gorm:"size:16;comment:请求方式;"`          //请求方式
	Permission   string `json:"permission" gorm:"size:128;comment:页面权限标识;"`   //页面权限标识
	PermissionId uint64 `json:"permission_id" gorm:"index;comment:菜单对应权限id;"` // 菜单对应权限id
	ParentId     uint64 `json:"parent_id" gorm:"size:11;comment:菜单父级id;"`     //菜单父级id
	NoCache      uint8  `json:"no_cache" gorm:"default:0;comment:缓存;"`        //缓存
	Component    string `json:"component" gorm:"size:255;comment:组件路径;"`      // 组件路径
	Sort         uint8  `json:"sort" gorm:"comment:菜单排序;"`                    // 排序
	Visible      uint8  `json:"visible" gorm:"comment:菜单状态;"`                 // 菜单状态
	CreateBy     string `json:"create_by" gorm:"size:128;comment:创建者;"`       // 创建者
	UpdateBy     string `json:"update_by" gorm:"size:128;comment:修改者;"`       // 修改者
	IsFrame      uint8  `json:"is_frame" gorm:"default:0;comment:是否外链;"`      // 是否外链
	Children     []Menu `json:"children" gorm:"-"`
	BaseModel
}

func (Menu) TableName() string {
	return "sys_menu"
}

func (e *Menu) SetMenuRole(roleId uint64) (m []Menu, err error) {

	menulist, err := e.GetByRole(roleId)

	m = make([]Menu, 0)
	for i := 0; i < len(menulist); i++ {
		if menulist[i].ParentId != 0 {
			continue
		}
		menusInfo := DiguiMenu(&menulist, menulist[i])

		m = append(m, menusInfo)
	}
	return
}

func (e *Menu) GetByRole(roleId uint64) (Menus []Menu, err error) {
	table := global.Eloquent.Table(e.TableName()).Select("sys_menu.*").Joins("left join sys_role_menu on sys_role_menu.menu_id=sys_menu.id")
	table = table.Where("sys_role_menu.role_id=? and menu_type in (?,?) and visible = ?", roleId, 1, 2, 0)
	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func DiguiMenu(menulist *[]Menu, menu Menu) Menu {
	list := *menulist

	min := make([]Menu, 0)
	for j := 0; j < len(list); j++ {

		if menu.ID != list[j].ParentId {
			continue
		}
		mi := Menu{}
		mi.ID = list[j].ID
		mi.Name = list[j].Name
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.Action = list[j].Action
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.NoCache = list[j].NoCache
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Visible = list[j].Visible
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []Menu{}

		if mi.MenuType != 3 {
			ms := DiguiMenu(menulist, mi)
			min = append(min, ms)

		} else {
			min = append(min, mi)
		}

	}
	menu.Children = min
	return menu
}
