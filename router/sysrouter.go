package router

import (
	"mime"

	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/api/v1/system"
	"github.com/isjyi/os/handler"
	"github.com/isjyi/os/pkg/jwt"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	sysStaticFileRouter(g)

	// swagger；注意：生产环境可以注释掉
	sysSwaggerRouter(g)

	// 无需认证
	sysNoCheckRoleRouter(g)
	// 需要认证
	// sysCheckRoleRouterInit(g, authMiddleware)

	return g
}

func sysBaseRouter(r *gin.RouterGroup) {

	// go ws.WebsocketManager.Start()
	// go ws.WebsocketManager.SendService()
	// go ws.WebsocketManager.SendAllService()

	r.GET("/", system.HelloWorld)

	// r.GET("/ws", ws.WebsocketManager.WsClient)

	r.GET("/info", handler.Ping)
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	mime.AddExtensionType(".js", "application/javascript")

	r.Static("/static", "./static")
	r.Static("/form-generator", "./static/form-generator")
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	url := ginSwagger.URL("http://localhost:8000/static/docs/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	v1.POST("/register", system.Register)
	v1.GET("/captcha", system.GenerateCaptchaHandler)
}
