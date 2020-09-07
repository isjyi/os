package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/isjyi/os/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	R = Router.Group("/base")
	{
		R.POST("register", v1.Register)
	}

	return
}
