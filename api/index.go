package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/pkg/redis"
	"go.uber.org/zap"
)

func Set(c *gin.Context) {
	name := c.DefaultQuery("name", "jerry")
	ok, err := redis.Exec("SET", "TEST", name)

	if err != nil {
		global.OS_LOG.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, gin.H{
		"msg": ok,
	})
}

func Get(c *gin.Context) {
	ok, err := redis.Exec("GET", "TEST")

	if err != nil {
		global.OS_LOG.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

	}

	c.JSON(200, gin.H{
		"msg": ok,
	})
}
