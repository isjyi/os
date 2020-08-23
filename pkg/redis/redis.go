package redis

import (
	"time"

	red "github.com/gomodule/redigo/redis"
)

type RedisCli struct {
	pool *red.Pool
}

func New() *RedisCli {
	return &RedisCli{
		pool: &red.Pool{
			MaxIdle:     256,
			MaxActive:   0,
			IdleTimeout: time.Duration(120),
			Dial: func() (red.Conn, error) {
				return red.Dial(
					"tcp",
					"127.0.0.1:6379",
					red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
					red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
					red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
					red.DialDatabase(0),
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
