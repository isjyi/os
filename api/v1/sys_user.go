package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global/response"
	"github.com/isjyi/os/resource"
	"github.com/isjyi/os/service"
	"github.com/isjyi/os/utils"
)

// 用户注册账号
// @Summary 用户注册账号
// @Description 用户注册接口
// @Tags Base
// @Accept  json
// @Produce  json
// @Param data body service.RegisterStruct true "param"
// @Success 200 {object} response.Response{data=resource.SysUserResource}"
// @Success 400 {object} response.Response"
// @Router /base/register [post]
func Register(c *gin.Context) {
	var R service.RegisterStruct
	if err := c.ShouldBindJSON(&R); err != nil {
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	u, err := R.Register()
	if err != nil {
		response.FailWithDetailed(response.ERROR, resource.SysUserResource{User: u}, fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(resource.SysUserResource{User: u}, "注册成功", c)
	}
}

func Login(c *gin.Context) {

}
