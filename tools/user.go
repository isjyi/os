package tools

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// func ExtractClaims(c *gin.Context) jwt.MapClaims {
// 	claims, exists := c.Get(jwt.JwtPayloadKey)
// 	if !exists {
// 		return make(jwt.MapClaims)
// 	}

// 	return claims.(jwt.MapClaims)
// }

func GetUserIdStr(c *gin.Context) string {
	val, exists := c.Get("userId")
	if !exists {
		fmt.Println(GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserIdStr 缺少UserId")
		return ""
	}
	return Int64ToString(int64((val).(float64)))
}

func GetUserIdUint64(c *gin.Context) uint64 {
	val, exists := c.Get("userId")
	if !exists {
		fmt.Println(GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetUserIdStr 缺少UserId")
		return 0
	}
	return uint64((val).(float64))
}
