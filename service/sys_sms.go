package service

import (
	"math/rand"
	"strings"
	"time"

	red "github.com/garyburd/redigo/redis"
	"github.com/isjyi/os/model/request"
	"github.com/isjyi/os/pkg/redis"
)

const (
	EX = 300
)

// @title    Captcha
// @description   获取验证码
// @auth                     （2020/09/10 11:33）
// @return    code            int
// @return    err       			error
func Captcha(s request.CaptchaStruct) (code int, err error) {
	var build strings.Builder

	build.WriteString("u:code:")
	build.WriteString(s.Phone)
	sname := build.String()

	code, err = redis.Int(redis.Exec("GET", sname))

	if err == nil {
		return
	}

	if err != nil {
		if err == red.ErrNil {
			err = nil
		} else {
			return
		}
	}

	rand.Seed(time.Now().Unix())

	code = rand.Intn(8000) + 1000

	redis.Exec("SET", sname, code, "EX", EX)

	// templateParam, err := global.OS_CONFIG.SMS.Marshal(code)

	// if err != nil {
	// 	return
	// }

	// if err = dysms.SendSms(global.OS_CONFIG.SMS.AccessKeyId, global.OS_CONFIG.SMS.AccessSecret, s.Phone, "颖橙密室管家", templateParam, "SMS_91700011"); err != nil {
	// 	return
	// }
	return
}
