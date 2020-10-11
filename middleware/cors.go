package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowMethods("OPTIONS")
	config.AllowOrigins = []string{"*"}
	config.AddAllowHeaders("AccessToken", "X-CSRF-Token", "Authorization",
		"Token", "X-Token", "X-User-Id")
	config.AddExposeHeaders("Content-Length", "Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers", "Content-Type")

	return cors.New(config)
}
