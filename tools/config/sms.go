package config

import "encoding/json"

type SMS struct {
	AccessKeyId  string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessSecret string `mapstructure:"access-secret" json:"accessSecret" yaml:"access-secret"`
}

func (s SMS) Marshal(code int) (param string, err error) {
	c := struct {
		Code int `json:"code"`
	}{
		code,
	}

	str, err := json.Marshal(c) //json序列化
	if err != nil {
		return
	}

	param = string(str[:])
	return
}
