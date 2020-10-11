package config

type Application struct {
	ReadTimeout    int    `mapstructure:"read-timeout" json:"readTimeout" yaml:"read-timeout"`
	WriterTimeout  int    `mapstructure:"writer-timeout" json:"writerTimeout" yaml:"writer-timeout"`
	EncryptionCost int    `mapstructure:"encryption-cost" json:"encryptionCost" yaml:"encryption-cost"`
	Host           string `mapstructure:"host" json:"host" yaml:"host"`
	Port           string `mapstructure:"port" json:"port" yaml:"port"`
	Name           string `mapstructure:"name" json:"name" yaml:"name"`
	Mode           string `mapstructure:"mode" json:"mode" yaml:"mode"`
}
