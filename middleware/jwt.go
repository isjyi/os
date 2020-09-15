package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中token为空",
			})
			c.Abort()
			return
		}

		claims, err := global.OS_JWT.ParseToken(authHeader)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + 60*60*24*7
			newToken, _ := global.OS_JWT.CreateToken(*claims)
			c.Header("Authorization", newToken)
		}

		c.Set("claims", claims)
		c.Next()
	}
}
