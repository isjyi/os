package config

import (
	"fmt"
)

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
}

type System struct {
	Env       string `mapstructure:"env" json:"env" yaml:"env"`
	Addr      int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	StaticDir string `mapstructure:"static_dir" json:"static_dir" yaml:"static_dir"`
	DbType    string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	Mode      string `mapstructure:"mode" json:"mode" yaml:"mode"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	MaxLifetime  int    `mapstructure:"max_life_time" json:"max_life_time" yaml:"max_life_time"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Network  string `mapstructure:"network" json:"network" yaml:"network"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type Log struct {
	File        string `mapstructure:"file" json:"file" yaml:"file"`
	Level       int    `mapstructure:"level" json:"level" yaml:"level"`
	MaxSize     int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxAge      int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress    bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
	ServiceName string `mapstructure:"service_name" json:"service_name" yaml:"service_name"`
}

func (db Mysql) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Addr,
		db.Dbname,
	)
}