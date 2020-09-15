package initialize

import (
	"os"
	"time"

	red "github.com/gomodule/redigo/redis"

	"github.com/garyburd/redigo/redis"
	"github.com/isjyi/os/global"
	"go.uber.org/zap"
)

func Redis() {
	c := global.OS_CONFIG.Redis

	global.OS_REDIS = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				c.Network,
				c.Addr,
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(c.DB),
			)
		},
	}

	conn := global.OS_REDIS.Get()

	defer conn.Close()

	if err := conn.Err(); err != nil {
		global.OS_LOG.Error("REDIS连接失败", zap.Error(err))
		os.Exit(0)
	}
	_, err := red.String(conn.Do("PING"))

	if err != nil {
		global.OS_LOG.Error(err.Error())
		os.Exit(0)
	}

	// global.OS_LOG.Info(fmt.Sprintf("redis connect ping response: %s", pong))
}
