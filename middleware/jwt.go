package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/pkg/jwt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(parts[1])

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		c.Set("id", claims.ID)
		c.Next()
	}
}
