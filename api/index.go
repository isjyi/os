package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/isjyi/os/pkg/check"
	"github.com/isjyi/os/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Set(c *gin.Context) {
	name := c.DefaultQuery("name", "jerry")
	ok, err := check.Rcli.Exec("SET", "TEST", name)

	if err != nil {
		log.Logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, gin.H{
		"msg": ok,
	})
}

func Get(c *gin.Context) {
	_, err := redis.String(check.Rcli.Exec("GET", "TEST"))

	if err != nil {
		log.Logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

	}

	c.JSON(200, gin.H{
		"msg": viper.GetString("test"),
	})
}
