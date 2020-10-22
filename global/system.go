package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CasbinEnforcer *casbin.SyncedEnforcer
	Eloquent       *gorm.DB
	Logger         *zap.Logger
	ReqLogger      *zap.Logger
	Version        string
	Check          *redis.Pool
	RedisPool      *redis.Pool
)

func init() {
	Version = "1.0.0"
}
