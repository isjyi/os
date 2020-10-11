package config

type Database struct {
	Driver       string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Source       string `mapstructure:"source" json:"source" yaml:"source"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	MaxLifetime  int    `mapstructure:"max-life-time" json:"maxLifeTime" yaml:"max-life-time"`
}
