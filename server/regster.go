package server

import (
	"errors"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/config"
	"github.com/jinzhu/gorm"
	"github.com/mojocn/base64Captcha"
)

// Register user register struct
type Register struct {
	Phone    string `json:"phone" label:"手机号" binding:"required,len=11,phone"`   // 手机号
	Password string `json:"passwrod" label:"密码" binding:"required,min=8,max=20"` // 密码
	NickName string `json:"nickname" binding:"omitempty,min=5,max=12"`           // 用户昵称
	VCode
} //@name Register
var store = base64Captcha.DefaultMemStore

func (reg *Register) Register() (err error) {
	var user models.SysUser

	// if !store.Verify(reg.UUID, reg.Code, true) {
	// 	return jwt.ErrInvalidVerificationode
	// }

	if !errors.Is(global.Eloquent.Where("phone = ?", reg.Phone).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册")
	}

	if p, err := tools.GenerateFromPassword(reg.Password, config.OSConfig.Application.EncryptionCost); err != nil {
		return err
	} else {
		reg.Password = p
	}

	user = models.SysUser{Phone: reg.Phone, NickName: reg.NickName, Password: reg.Password}

	err = global.Eloquent.Create(&user).Error

	return
}
