package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/pkg/jwt"
	"github.com/isjyi/os/server"
	"github.com/isjyi/os/tools/config"
	"github.com/isjyi/os/utils"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(models.SysUser)
		r, _ := v["role"].(models.SysRole)
		return jwt.MapClaims{
			jwt.IdentityKey: u.ID,
			jwt.RoleIdKey:   strconv.Itoa(int(r.ID)),
			jwt.NiceKey:     u.NickName,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims[jwt.IdentityKey],
		"NickName":    claims[jwt.NiceKey],
		"RoleIds":     claims[jwt.RoleIdKey],
	}
}

// @Summary 登陆
// @Description 获取token
// @Description LoginHandler can be used by clients to get a jwt token.
// @Description Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// @Description Reply will be of the form {"token": "TOKEN"}.
// @Description dev mode：It should be noted that all fields cannot be empty, and a value of 0 can be passed in addition to the account password
// @Description 注意：开发模式：需要注意全部字段不能为空，账号密码外可以传入0值
// @Tags Base
// @Accept  application/json
// @Product application/json
// @Param account body server.Login  true "account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /v1/login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals server.Login

	if err := c.ShouldBindJSON(&loginVals); err != nil {
		return nil, err
	}
	if config.OSConfig.Application.Mode != "dev" {
		if !store.Verify(loginVals.UUID, loginVals.Code, true) {
			return nil, jwt.ErrInvalidVerificationode
		}
	}
	user, role, e := loginVals.GetUser()
	if e == nil {
		return map[string]interface{}{"user": user, "role": role}, nil
	} else {
		global.ReqLogger.Error(e.Error())
	}

	return nil, jwt.ErrFailedAuthentication
}

// @Summary 退出登录
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Reply will be of the form {"token": "TOKEN"}.
// @Tags Base
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "msg": "成功退出系统" }"
// @Router /v1/logout [post]
// @Security Bearer
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}

func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		c.Set("roleId", v["RoleIds"])
		c.Set("userId", v["IdentityKey"])
		c.Set("nickname", v["NickName"])
		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(code, gin.H{
			"code": code,
			"msg":  "数据验证非法",
			"data": utils.T(errs),
		})
	}

}
