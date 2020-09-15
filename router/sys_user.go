package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/isjyi/os/api/v1"
	"github.com/isjyi/os/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.JWTAuthMiddleware())
	{
		UserRouter.POST("me", v1.Me) // 获取用户信息
	}
}
