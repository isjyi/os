package middleware

import (
	"time"

	"github.com/isjyi/os/handler"
	"github.com/isjyi/os/pkg/jwt"
	"github.com/isjyi/os/tools/config"
)

func AuthInit() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "os",
		Key:             []byte(config.OSConfig.Jwt.Secret),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		PayloadFunc:     handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler,
		Authenticator:   handler.Authenticator,
		Authorizator:    handler.Authorizator,
		Unauthorized:    handler.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}
