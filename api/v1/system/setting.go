package system

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
)

// @Summary 查询系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/setting [get]
func GetSetting(c *gin.Context) {
	var s models.SysSetting

	r, e := s.Get()

	tools.HasError(e, "查询失败", 500)

	app.OK(c, r, "查询成功")
}
