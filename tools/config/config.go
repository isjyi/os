package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Application Application `mapstructure:"application" json:"application" yaml:"application"`
	Jwt         Jwt         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Database    Database    `mapstructure:"database" json:"database" yaml:"database"`
	Redis       Redis       `mapstructure:"redis" json:"redis" yaml:"redis"`
	Logger      Logger      `mapstructure:"logger" json:"logger" yaml:"logger"`
	SMS         SMS         `mapstructure:"sms" json:"sms" yaml:"sms"`
}

var (
	OSConfig *Config
)

// 载入配置文件
func Setup(path string) {
	v := viper.New()

	v.SetConfigFile(path)

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	v.WatchConfig()

	v.OnConfigChange(func(event fsnotify.Event) {
		fmt.Println("config file changed:", event.Name)
		if err := v.Unmarshal(&OSConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&OSConfig); err != nil {
		fmt.Println(err)
	}
}
