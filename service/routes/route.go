package route

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/api"
	"github.com/isjyi/os/service/middleware"
)

func Routes(r *gin.Engine) {

	r.GET("/set", api.Set)

	r.GET("/get", api.Get)

	authed := r.Group("/")

	middleware.Auth(authed)

}
