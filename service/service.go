package service

import (
	"fmt"

	"github.com/isjyi/os/pkg/log"
	"github.com/isjyi/os/pkg/mysql"
	"github.com/isjyi/os/pkg/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type application struct {
	logger *zap.Logger
	db     *gorm.DB
	redis  *redis.RedisCli
}

func ListenAndServe() {
	db, err := mysql.New(log.Logger)

	if err != nil {
		log.Logger.Error(err.Error(), zap.Error(err))
		return
	}

	defer db.Close()

	app := &application{
		logger: log.Logger,
		db:     db,
		redis:  redis.New(),
	}
	fmt.Println(app)
}
