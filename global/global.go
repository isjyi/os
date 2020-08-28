package global

import (
	"github.com/garyburd/redigo/redis"
	"github.com/isjyi/os/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	OS_DB     *gorm.DB
	OS_REDIS  *redis.Pool
	OS_CONFIG config.Server
	OS_VP     *viper.Viper
	OS_LOG    *zap.Logger
)
