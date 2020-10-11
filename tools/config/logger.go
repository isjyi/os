package config

type Logger struct {
	Path       string `mapstructure:"path" json:"path" yaml:"path"`
	Level      int    `mapstructure:"level" json:"level" yaml:"level"`
	MaxSize    int    `mapstructure:"max-size" json:"maxSize" yaml:"max-size"`
	MaxBackups int    `mapstructure:"max-backups" json:"maxBackups" yaml:"max-backups"`
	MaxAge     int    `mapstructure:"max-age" json:"maxAge" yaml:"max-age"`
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
	Stdout     bool   `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	EnabledBUS bool   `mapstructure:"enabledbus" json:"enabledbus" yaml:"enabledbus"`
	EnabledREQ bool   `mapstructure:"enabledreq" json:"enabledreq" yaml:"enabledreq"`
	EnabledDB  bool   `mapstructure:"enableddb" json:"enableddb" yaml:"enableddb"`
	EnabledJOB bool   `mapstructure:"enabledjob" json:"enabledjob" yaml:"enabledjob"`
}
