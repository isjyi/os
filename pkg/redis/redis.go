package redis

import (
	"github.com/isjyi/os/global"
)

func Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := global.OS_REDIS.Get()
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
