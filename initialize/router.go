package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/api"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/middleware"
	"github.com/isjyi/os/router"
	_ "github.com/isjyi/os/swagger"
	"github.com/rakyll/statik/fs"
)

func Routers() *gin.Engine {
	gin.SetMode(global.OS_CONFIG.System.Mode)

	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// global.OS_LOG.Debug("use middleware logger")

	// 跨域
	Router.Use(middleware.Cors())
	global.OS_LOG.Info("use middleware cors")

	if global.OS_CONFIG.System.Mode == "debug" {
		statikFS, err := fs.New()
		if err != nil {
			panic(err)
		}
		Router.StaticFS("/swagger", statikFS)
		global.OS_LOG.Info("swagger register success")
	}
	v1 := Router.Group("/api/v1")
	Router.GET("/set", api.Set)

	Router.GET("/get", api.Get)

	router.InitBaseRouter(v1)
	router.InitUserRouter(v1)

	global.OS_LOG.Info("router register success")

	return Router
}
