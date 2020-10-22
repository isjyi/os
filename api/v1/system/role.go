package system

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
)

// @Summary 获取Role数据
// @Description 获取JSON
// @Tags Base
// @Param roleId path string false "roleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /v1/role [get]
// @Security Bearer
func GetRole(c *gin.Context) {
	var role models.SysRole
	role.ID, _ = tools.StringToUInt64(c.GetString("roleId"))
	err := global.Eloquent.Where("id = ? ", role.ID).First(&role).Error

	if err != nil {
		tools.HasError(err, "抱歉未找到相关信息", -1)
	}

	app.OK(c, role, "")
}
