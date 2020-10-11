package captcha

import (
	"log"

	"github.com/isjyi/os/pkg/redis"
	"github.com/mojocn/base64Captcha"
)

type redisStore struct {
	// Expiration time of captchas.
	expiration int
}

func NewRedisStore(expiration int) base64Captcha.Store {
	s := new(redisStore)
	s.expiration = expiration
	return s
}

func (s *redisStore) Set(id string, value string) {
	_, err := redis.Exec("SETEX", id, s.expiration, value)

	if err != nil {
		log.Println(err.Error())
	}
}

func (s *redisStore) Get(id string, clear bool) (value string) {

	val, err := redis.String(redis.Exec("GET", id))
	if err != nil {
		log.Println(err.Error())
		return
	}

	if clear {
		_, err := redis.Exec("DEL", id)
		if err != nil {
			log.Println(err.Error())
			return ""
		}
	}
	return val
}

func (s *redisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
