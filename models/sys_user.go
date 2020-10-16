package models

type SysUser struct {
	ID        uint64 `gorm:"primary_key"`                                                                                // 用户id
	Phone     string `gorm:"size:11;unique_index;not null;comment:'用户登录手机号'" json:"phone"`                               // 手机号
	Password  string `json:"-"  gorm:"size:64;comment:'用户登录密码'"`                                                         // 密码
	NickName  string `json:"nickName" gorm:"size:20;default:'jerry';comment:'用户昵称'" `                                    // 昵称
	HeaderImg string `json:"headerImg" gorm:"size:128;default:'http://qmplusimg.henrongyi.top/head.png';comment:'用户头像'"` //头像
	RoleId    uint64 `gorm:"comment:'角色id';index" json:"role_id"`                                                        // 角色id
	LogindAt  int    `json:"-" gorm:"default:null"`                                                                      // 登录时间
	BaseModel
}

func (SysUser) TableName() string {
	return "sys_user"
}
