package middleware

import (
	"fmt"
	"net/http"

	mycasbin "github.com/isjyi/os/pkg/casbin"
	"github.com/isjyi/os/pkg/jwt"
	"github.com/isjyi/os/tools"

	"github.com/gin-gonic/gin"
)

//权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("JWT_PAYLOAD")
		v := data.(jwt.MapClaims)
		e, err := mycasbin.Casbin()
		tools.HasError(err, "", 500)
		//检查权限
		res, err := e.Enforce(v["rid"], c.Request.URL.Path, c.Request.Method)
		tools.HasError(err, "", 500)

		fmt.Printf("%s [INFO] %s %s %s \r\n",
			tools.GetCurrentTimeStr(),
			c.Request.Method,
			c.Request.URL.Path,
			v["rid"],
		)

		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}
