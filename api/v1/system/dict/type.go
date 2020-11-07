package dict

import (
	"github.com/gin-gonic/gin"
	"github.com/isjyi/os/models"
	"github.com/isjyi/os/tools"
	"github.com/isjyi/os/tools/app"
)

// @Summary 字典类型列表数据
// @Description 获取JSON
// @Tags 字典类型
// @Param name query string false "字典类型名称"
// @Param id query string false "字典id""
// @Param type query string false "字典类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.PageResponse "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type/list [get]
// @Security Bearer
func GetDictTypeList(c *gin.Context) {
	var data models.SysDictType
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = tools.StringToInt(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = tools.StringToInt(index)
	}

	data.Name = c.Request.FormValue("name")
	id := c.Request.FormValue("id")
	data.Id, _ = tools.StringToUInt64(id)
	data.Type = c.Request.FormValue("type")
	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}
