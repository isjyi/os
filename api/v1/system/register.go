package system

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/server"
	"github.com/isjyi/os/tools/app"
)

// @Summary 用户注册账号
// @Description 用户注册接口
// @Tags Base
// @Accept  json
// @Produce  json
// @Param data body server.Register true "param"
// @Success 200 {object} app.Response"
// @Failure 402 {object} app.Response"
// @Router /api/v1/register [post]
func Register(c *gin.Context) {
	var r server.Register
	if err := c.ShouldBindJSON(&r); err != nil {
		app.Error(c, 402, err, "数据验证失败")
		return
	}

	err := r.Register()
	if err != nil {
		app.Error(c, 402, err, "")
	} else {
		app.OK(c, "", "注册成功")
	}
}
