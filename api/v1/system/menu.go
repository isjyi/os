package system

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/global"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
	"github.com/isjyi/os/tools/app/msg"
)

// @Summary 根据角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {object} app.Response{data=[]models.Menu} "desc"
// @Failure 401 {object} app.Response"
// @Router /api/v1/menurole [get]
// @Security Bearer
func MenuRole(c *gin.Context) {
	var user models.SysUser

	if err := global.Eloquent.Preload("Role").First(&user, tools.GetUserIdStr(c)).Error; err != nil {
		tools.HasError(err, msg.UserFound, 401)
	}

	var menu models.Menu

	result, err := menu.SetMenuRole(user.Role.ID)

	if err != nil {
		app.Error(c, -1, err, "获取失败")
		return
	}

	app.OK(c, result, "")
}
