package core

import (
	"fmt"
	"time"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/initialize"
	"github.com/isjyi/os/utils"
)

type server interface {
	ListenAndServe() error
}

func ListenAndServe() {

	defer global.OS_DB.Close()

	Router := initialize.Routers()

	Router.Static("/public", global.OS_CONFIG.System.StaticDir)

	address := fmt.Sprintf(":%d", global.OS_CONFIG.System.Addr)

	utils.InitTrans("zh")

	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)

	global.OS_LOG.Debug(fmt.Sprintf("server run success on %s", address))

	err := s.ListenAndServe()
	if err != nil {
		global.OS_LOG.Error(err.Error())
	}
}
