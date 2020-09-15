package initialize

import (
	"time"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/pkg/jwt"
)

func JWTManager() {
	c := global.OS_CONFIG.JWT
	global.OS_JWT = &jwt.JWTManager{SecretKey: c.SigningKey, ExpiresAt: time.Hour * time.Duration(c.ExpiresAt)}
}
