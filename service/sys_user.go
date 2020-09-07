package service

import (
	"errors"
	"fmt"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/model"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// user register struct
type RegisterStruct struct {
	Phone    string `json:"phone" label:"手机号" binding:"required,len=11"`
	Password string `json:"passwrod" label:"密码" binding:"required,min=8,max=20"`
	NickName string `json:"nickname" binding:"omitempty,min=6,max=12"`
}

func (s *RegisterStruct) Register() (user model.User, err error) {
	if !errors.Is(global.OS_DB.Where("phone = ?", s.Phone).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return user, errors.New("用户名已注册")
	}

	user = model.User{Phone: s.Phone, NickName: s.NickName, Password: s.Password}
	user.SetPassword(s.Password)
	fmt.Println(user.Password)
	user.UUID = uuid.NewV4()
	err = global.OS_DB.Create(&user).Error

	return
}

type LoginStruct struct {
	Phone    string `json:"phone" label:"手机号" binding:"required,len=11"`
	Password string `json:"passwrod" label:"密码" binding:"required,min=8,max=20"`
}

// func (s *LoginStruct) Login ()
