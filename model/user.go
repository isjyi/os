package model

import (
	"time"

	"github.com/isjyi/os/global"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint       `gorm:"primary_key"`
	UUID      uuid.UUID  `json:"uuid" gorm:"comment:'用户UUID'"`
	Phone     string     `gorm:"size:11;unique;not null;index:idx_users_phone" json:"phone"`
	Password  string     `json:"-"  gorm:"size:64;comment:'用户登录密码'"`
	NickName  string     `json:"nickName" gorm:"size:20;default:'jerry';comment:'用户昵称'" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), global.OS_CONFIG.System.EncryptionCost)
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
