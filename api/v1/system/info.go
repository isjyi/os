package system

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
)

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags Base
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /v1/info [get]
// @Security Bearer
func Info(c *gin.Context) {
	var user models.SysUser

	if err := global.Eloquent.Preload("Role").First(&user, tools.GetUserIdStr(c)).Error; err != nil {
		tools.HasError(err, "", 500)
	}

	var mp = make(map[string]interface{})

	mp["roles"] = []string{user.Role.RoleName}

	if user.Role.RoleName == "系统管理员" {
		mp["permissions"] = []string{"*:*:*"}
	} else {
		RoleMenu := models.SysRoleMenu{RoleId: user.RoleId}
		list, _ := RoleMenu.GetPermits()
		mp["permissions"] = list
	}

	mp["introduction"] = " am a super administrator"

	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if user.Avatar != "" {
		mp["avatar"] = user.Avatar
	}
	mp["userName"] = user.Phone
	mp["userId"] = user.ID
	mp["name"] = user.NickName
	app.OK(c, mp, "")
}
