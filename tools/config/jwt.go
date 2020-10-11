package config

type Jwt struct {
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Timeout int64  `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}
