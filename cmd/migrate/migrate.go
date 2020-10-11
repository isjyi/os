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
package migrate

import (
	"fmt"

	"github.com/isjyi/os/database"
	"github.com/isjyi/os/database/migrate"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/pkg/logger"
	"github.com/isjyi/os/tools/config"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var (
	configYml  string
	MigrateCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "os migrate -c config/os.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	MigrateCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/os.yaml", "Start server with provided configuration file")
}

func run() {
	usage := `start init`
	fmt.Println(usage)

	//1. 读取配置
	config.Setup(configYml)
	//2. 设置日志
	logger.Setup()
	//3. 初始化数据库链接
	database.Setup(config.OSConfig.Database.Driver)
	//4. 数据库迁移
	_ = migrateModel()
	//5. 数据初始化完成
	// if err := models.InitDb(); err != nil {
	// 	global.Logger.Fatal("数据库基础数据初始化失败！")
	// }

	usage = `数据库基础数据初始化成功`
	fmt.Println(usage)
}

func migrateModel() error {
	if config.OSConfig.Database.Driver == "mysql" {
		global.Eloquent = global.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	return migrate.AutoMigrate(global.Eloquent)
}
