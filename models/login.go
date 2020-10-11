package models

import (
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/tools"
)

type Login struct {
	Phone    string `form:"Phone" json:"phone" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

func (u *Login) GetUser() (user SysUser, role SysRole, e error) {

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
