package redis

import (
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	red "github.com/garyburd/redigo/redis"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/config"
	"go.uber.org/zap"
)

/**
 * 实例化
 */
func Setup() {
	global.Check = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				config.OSConfig.Redis.Network,
				config.OSConfig.Redis.Addr,
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(config.OSConfig.Redis.DB),
			)
		},
	}

	conn := global.Check.Get()

	defer conn.Close()

	if err := conn.Err(); err != nil {
		global.Logger.Sugar().Error("REDIS连接失败", zap.Error(err))
		os.Exit(0)
	}
	_, err := red.String(conn.Do("PING"))

	if err != nil {
		global.Logger.Sugar().Error(err.Error())
		os.Exit(0)
	}
	global.Logger.Info(tools.Green("redis connect success !"))
}

func Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := global.Check.Get()
	if err := con.Err(); err != nil {

	}

	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}

func String(reply interface{}, err error) (string, error) {
	return red.String(reply, err)
}

func Int(reply interface{}, err error) (int, error) {
	return red.Int(reply, err)
}
