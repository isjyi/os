package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	CasbinEnforcer *casbin.SyncedEnforcer
	Eloquent       *gorm.DB
	Logger         *zap.Logger
	ReqLogger      *zap.Logger
	Version        string
	RedisPool      *redis.Pool
)

func init() {
	Version = "1.0.0"
}
