/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/database"
	"github.com/isjyi/os/global"
	mycasbin "github.com/isjyi/os/pkg/casbin"
	"github.com/isjyi/os/pkg/logger"
	"github.com/isjyi/os/pkg/redis"
	"github.com/isjyi/os/router"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/config"
	"github.com/isjyi/os/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serverCmd represents the server command
var (
	configYml string
	port      string
	mode      string
	ServerCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "os server -c config/os.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	ServerCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/os.yaml", "Start server with provided configuration file")
}

func setup() {

	//1. 读取配置
	config.Setup(configYml)
	//2. 设置日志
	logger.Setup()
	//3. 初始化数据库链接
	database.Setup(config.OSConfig.Database.Driver)
	//4. 初始化redis链接
	redis.Setup()
	//5. 接口访问控制加载
	mycasbin.Setup()

	usageStr := `starting api server`
	global.Logger.Info(usageStr)

}

func run() error {

	if config.OSConfig.Application.Mode == string(tools.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()

	defer global.Eloquent.Close()

	utils.InitTrans("zh")

	srv := &http.Server{
		Addr:    config.OSConfig.Application.Host + ":" + config.OSConfig.Application.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Fatal("listen: ", zap.Error(err))
		}
	}()
	tip()
	fmt.Println(tools.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", config.OSConfig.Application.Port)
	fmt.Printf("-  Network: http://%s:%s/ \r\n", tools.GetLocaHonst(), config.OSConfig.Application.Port)
	fmt.Println(tools.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/swagger/index.html \r\n", config.OSConfig.Application.Port)
	fmt.Printf("-  Network: http://%s:%s/swagger/index.html \r\n", tools.GetLocaHonst(), config.OSConfig.Application.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", tools.GetCurrentTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Fatal("Server Shutdown:", zap.Error(err))
	}
	fmt.Printf("Server exiting")
	return nil
}

func tip() {
	content, _ := ioutil.ReadFile("./static/os.txt")
	fmt.Println(tools.Red(string(content)))
	usageStr := `欢迎使用 ` + tools.Green(`os `+global.Version) + ` 可以使用 ` + tools.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}
