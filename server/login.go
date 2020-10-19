package server

import (
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/tools"
)

// Login user login struct
type Login struct {
	Phone    string `json:"phone"  binding:"required,len=11,phone"`    // 手机号
	Password string `json:"passwrod"  binding:"required,min=8,max=20"` // 密码
	VCode
} //@name Login

// VCode user verification code
type VCode struct {
	Code string `json:"code"  binding:"required,len=5"` // 验证码
	UUID string `json:"uuid"  binding:"required"`       // 验证码唯一id
} //@name verificationCode

func (u *Login) GetUser() (user models.SysUser, role models.SysRole, e error) {

	e = global.Eloquent.Where("phone = ?", u.Phone).First(&user).Error
	if e != nil {
		return
	}
	_, e = tools.CompareHashAndPassword(user.Password, u.Password)
	if e != nil {
		return
	}
	e = global.Eloquent.Where("id = ? ", user.RoleId).First(&role).Error
	if e != nil {
		return
	}
	return
}
