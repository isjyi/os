package resource

import "github.com/isjyi/os/model"

type SysUserResource struct {
	User model.User `json:"user"`
}

type SysLoginResource struct {
	Token string `json:"token"`
}
