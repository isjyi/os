package model

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint         `gorm:"primary_key"`                                                                       // 用户id
	Phone     string       `gorm:"size:11;unique;not null;index:idx_users_phone;comment:'用户登录手机号'" json:"phone"`      // 手机号
	Password  string       `json:"-"  gorm:"size:64;comment:'用户登录密码'"`                                                // 密码
	NickName  string       `json:"nickName" gorm:"size:20;default:'jerry';comment:'用户昵称'" `                           // 昵称
	HeaderImg string       `json:"headerImg" gorm:"default:'http://qmplusimg.henrongyi.top/head.png';comment:'用户头像'"` //头像
	LogindAt  sql.NullTime `json:"-"`                                                                                 // 登录时间
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt *time.Time   `json:"-" sql:"index"`
}

func (u *User) SetPassword(password string, cost int) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}

	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
