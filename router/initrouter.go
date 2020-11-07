package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/middleware"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine

	r = gin.New()

	middleware.InitMiddleware(r)
	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// tools.HasError(err, "JWT Init Error", 500)

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	// 注册业务路由
	// TODO: 这里可存放业务路由，里边并无实际路由只有演示代码
	InitExamplesRouter(r, authMiddleware)
	registerDictRouter(v1, authMiddleware)
	return r
}
