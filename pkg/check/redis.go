package check

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

type RedisCli struct {
	pool *redis.Pool
}

var Rcli *RedisCli

func New() {
	Rcli = &RedisCli{
		pool: &redis.Pool{
			MaxIdle:     256,
			MaxActive:   0,
			IdleTimeout: time.Duration(120),
			Dial: func() (redis.Conn, error) {
				return redis.Dial(
					viper.GetString("redis.network"),
					viper.GetString("redis.address"),
					redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
					redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
					redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
					redis.DialDatabase(0),
				)
			},
		},
	}
}

func (r *RedisCli) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := r.pool.Get()
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
