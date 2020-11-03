package models

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
	NoCache      bool   `json:"no_cache" gorm:"default:0;comment:缓存;"`        //缓存
	Component    string `json:"component" gorm:"size:255;comment:组件路径;"`      // 组件路径
	Sort         uint8  `json:"sort" gorm:"comment:菜单排序;"`                    // 排序
	Visible      bool   `json:"visible" gorm:"comment:菜单状态;"`                 // 菜单状态
	CreateBy     string `json:"create_by" gorm:"size:128;comment:创建者;"`       // 创建者
	UpdateBy     string `json:"update_by" gorm:"size:128;comment:修改者;"`       // 修改者
	IsFrame      bool   `json:"is_frame" gorm:"default:0;comment:是否外链;"`      // 是否外链
	Children     []Menu `json:"children" gorm:"-"`
	BaseModel
}

func (Menu) TableName() string {
	return "sys_menu"
}
