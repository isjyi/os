package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global/response"
	"github.com/isjyi/os/model/request"
	"github.com/isjyi/os/pkg/jwt"
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
// @Param data body request.RegisterStruct true "param"
// @Success 200 {object} response.Response{data=resource.SysUserResource}"
// @Failure 400 {object} response.Response"
// @Router /base/register [post]
func Register(c *gin.Context) {
	var r request.RegisterStruct
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithDetailed(utils.ErrResp(err), c)
		return
	}

	u, err := service.Register(r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkDetailed(resource.SysUserResource{User: u}, "注册成功", c)
	}
}

// 用户登录
// @Summary 用户登录
// @Description 用户登录接口
// @Tags Base
// @Accept  json
// @Produce  json
// @Param data body request.LoginStruct true "param"
// @Success 200 {object} response.Response{data=resource.SysLoginResource}"
// @Failure 400 {object} response.Response"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var r request.LoginStruct
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithDetailed(utils.ErrResp(err), c)
		return
	}

	token, err := service.Login(r)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkDetailed(resource.SysLoginResource{Token: token}, "ok", c)
	}
}

// 获取用户信息
// @Summary 用户信息
// @Description 获取用户信息接口
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=resource.SysUserResource}"
// @Failure 400 {object} response.Response"+
// @Security JWT
// @Router /user/me [post]
func Me(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*jwt.UserClaims)

	user, err := service.Info(waitUse.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkDetailed(resource.SysUserResource{User: user}, "ok", c)
	}
}
