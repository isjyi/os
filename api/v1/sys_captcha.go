package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global/response"
	"github.com/isjyi/os/model/request"
	"github.com/isjyi/os/service"
	"github.com/isjyi/os/utils"
)

// @Summary 发送验证码
// @Tags Base
// @Accept  json
// @Produce  json
// @Param data body request.CaptchaStruct true "param"
// @Success 200 {object} response.Response"
// @Failure 400 {object} response.Response"
// @Failure 500 {object} response.Response"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	var r request.CaptchaStruct
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithDetailed(utils.ErrResp(err), c)
		return
	}

	code, err := service.Captcha(r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkDetailed(code, "操作成功", c)
	}
}
