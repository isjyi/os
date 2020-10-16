package system

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
	"github.com/isjyi/os/tools/captcha"
)

// @Summary 获取验证码
// @Description 获取验证码
// @Tags Base
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"{"code": 200,"data": [...]}"
// @Router /api/v1/captcha [get]
func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	tools.HasError(err, "验证码获取失败", 500)
	app.Custum(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
