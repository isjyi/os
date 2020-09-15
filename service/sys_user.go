package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/model"
	"github.com/isjyi/os/model/request"
	"github.com/isjyi/os/pkg/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// @title    Register
// @description   register, 用户注册
// @auth                     2020/09/08 09:47
// @return    user            mode.User
// @return    err       			error
func Register(s request.RegisterStruct) (user model.User, err error) {
	if !errors.Is(global.OS_DB.Where("phone = ?", s.Phone).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return user, errors.New("用户名已注册")
	}

	var build strings.Builder

	build.WriteString("u:code:")
	build.WriteString(s.Phone)
	rname := build.String()

	rcode, err := redis.Int(redis.Exec("GET", rname))

	if err != nil {
		global.OS_LOG.Error(err.Error(), zap.Error(err))
		return user, errors.New("验证码不存在或已过期")
	}

	if s.Code != strconv.Itoa(rcode) {
		return user, errors.New("验证码不正确")
	}

	if _, e := redis.Exec("DEL", rname); e != nil {
		return user, e
	}

	user = model.User{Phone: s.Phone, NickName: s.NickName, Password: s.Password}
	user.SetPassword(s.Password, global.OS_CONFIG.System.EncryptionCost)
	err = global.OS_DB.Create(&user).Error

	return
}

// @title    Login
// @description   login, 用户登录
// @auth                     2020/9/10 17:8:58
// @return    token            string
// @return    err       			error
func Login(s request.LoginStruct) (token string, err error) {
	var user model.User

	if errors.Is(global.OS_DB.Where("phone = ?", s.Phone).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return token, errors.New("用户不存在")
	}

	if user.CheckPassword(s.Password) {
		token, err = global.OS_JWT.Generate(user)
		return
	}

	return token, errors.New("密码错误")
}

// @title    Info
// @description    获取用户信息
// @auth                     2020/09/14 09:47
// @return    user            mode.User
// @return    err       			error
func Info(id uint) (user model.User, err error) {
	err = global.OS_DB.Where("id = ?", id).First(&user).Error
	return
}
