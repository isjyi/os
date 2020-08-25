package service

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/pkg/check"
	"github.com/isjyi/os/pkg/db"
	"github.com/isjyi/os/pkg/log"
	"github.com/isjyi/os/service/middleware"
	route "github.com/isjyi/os/service/routes"
	"go.uber.org/zap"
)

func ListenAndServe() {
	log.New()

	db, err := db.New()

	if err != nil {
		log.Logger.Error(err.Error(), zap.Error(err))
		return
	}

	defer db.Close()

	check.New()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	middleware.Global(r)
	route.Routes(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
