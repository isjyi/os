package system

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/response"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
)

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 系统
// @Success 200 {object} app.Response{data=response.InfoResponse}
// @Router /v1/info [get]
// @Security Bearer
func Info(c *gin.Context) {
	var user models.SysUser

	if err := global.Eloquent.Preload("Role").First(&user, tools.GetUserIdStr(c)).Error; err != nil {
		tools.HasError(err, "", 500)
	}

	var res = response.InfoResponse{
		UserName: user.Phone,
		UserID:   user.ID,
		Name:     user.NickName,
		Roles:    []string{user.Role.RoleName},
	}

	if user.Role.RoleName == "管理员" {
		res.Permissions = []string{"*:*:*"}
	} else {
		RoleMenu := models.SysRoleMenu{RoleId: user.RoleId}
		list, _ := RoleMenu.GetPermits()
		res.Permissions = list
	}

	res.Aatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if user.Avatar != "" {
		res.Aatar = user.Avatar
	}
	app.OK(c, res, "")
}
